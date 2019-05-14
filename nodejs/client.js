const grpc = require("grpc");
const PROTO_PATH = "../proto/CoinPriceChart.proto";
const CoinPriceChart = grpc.load(PROTO_PATH).CoinPriceChart;
const client = new CoinPriceChart(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

client.PriceChart({ coin: "bitcoin" }, (error, notes) => {
  if (!error) {
    console.log("successfully fetch List notes");
    console.log(notes);
  } else {
    console.error(error);
  }
});
