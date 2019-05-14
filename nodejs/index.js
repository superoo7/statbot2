const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const captureWebsite = require("capture-website");
const AWS = require("aws-sdk");
const dotenv = require("dotenv");
const path = require("path");
const fs = require("fs");

// ========== CONFIG ==========
// DOTENV config
dotenv.config({
  path: path.resolve(process.cwd(), "../.env")
});

// AWS config
AWS.config.update({
  accessKeyId: process.env.ACCESS_KEY_ID,
  secretAccessKey: process.env.SECRET_ACCESS_KEY
});
const s3 = new AWS.S3();

function s3Upload(type, filePath) {
  const params = {
    Bucket: "statbot.superoo7.com",
    Body: fs.createReadStream(filePath),
    Key: `${type}/${path.basename(filePath)}`,
    ContentType: "image/png",
    ACL: "public-read"
  };

  return new Promise((resolve, reject) => {
    s3.upload(params, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
}

// GRPC config
const packageDefinition = protoLoader.loadSync(
  "../proto/CoinPriceChart.proto",
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }
);
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const CoinPriceChart = protoDescriptor.CoinPriceChart;

// ========== HELPER ==========

function log(message) {
  console.log(`[${new Date().toISOString()}]`, message);
}

// ========== GRPC SERVER ==========

const server = new grpc.Server();

server.addService(CoinPriceChart.service, {
  PriceChart: (data, callback) => {
    const { coin } = data.request;
    const timeNow = Date.now();
    const link = `https://superoo7.github.io/ce-to-iframe?wc-name=coingecko-coin-price-chart-widget&wc-src=https%3A%2F%2Fwidgets.coingecko.com%2Fcoingecko-coin-price-chart-widget.js&width=400&height=300&coin-id=${coin}`;
    const img = `../img/${coin}-${timeNow}.png`;
    captureWebsite
      .file(link, img, { height: 320, width: 420, delay: 2 })
      .then(() => {
        return s3Upload("PriceChart", img);
      })
      .then(data => {
        log(data.Location);
        callback(null, {
          fileName: `${coin}.png`,
          timestamp: timeNow,
          key: data.key
        });
      })
      .catch(err => {
        callback({
          code: 400,
          message: err.message,
          status: grpc.status.INTERNAL
        });
      });
  }
});

server.bind("127.0.0.1:50051", grpc.ServerCredentials.createInsecure());
console.log("Server running at http://127.0.0.1:50051");
server.start();
