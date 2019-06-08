package sh

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ararog/timeago"
	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
	"github.com/superoo7/statbot2/http"
)

type SteemHuntPost struct {
	ID      int      `json:"id"`
	Author  string   `json:"author"`
	URL     string   `json:"url"`
	Title   string   `json:"title"`
	Tagline string   `json:"tagline"`
	Tags    []string `json:"tags"`
	Images  []struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"images"`
	Beneficiaries []interface{} `json:"beneficiaries"`
	Permlink      string        `json:"permlink"`
	IsActive      bool          `json:"is_active"`
	PayoutValue   float64       `json:"payout_value"`
	ActiveVotes   []struct {
		Voter      string `json:"voter"`
		Weight     int    `json:"weight"`
		Rshares    string `json:"rshares"`
		Percent    int    `json:"percent"`
		Reputation int    `json:"reputation"`
		Time       string `json:"time"`
	} `json:"active_votes"`
	Children    int       `json:"children"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	IsVerified  bool      `json:"is_verified"`
	VerifiedBy  string    `json:"verified_by"`
	HuntScore   float64   `json:"hunt_score"`
	ValidVotes  []struct {
		Score   float64 `json:"score"`
		Voter   string  `json:"voter"`
		Weight  int     `json:"weight"`
		Percent int     `json:"percent"`
	} `json:"valid_votes"`
	ListedAt time.Time `json:"listed_at"`
}

func HuntCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage, link string) {
	// delete message
	d.Discord.ChannelMessageDelete(m.ChannelID, m.ID)

	authorName := ""
	permlinkName := ""
	match, err := regexp.MatchString("(https?:\\/\\/[^\\s]+)", link)
	if match == false || err != nil {
		em := d.GenErrorMessage("Not a valid link")
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}
	re := regexp.MustCompile("[\\/#]")
	splittedLink := re.Split(link, -1)
	permlinkCounter := 0

	for i := range splittedLink {
		l := splittedLink[i]
		if permlinkCounter != 0 {
			permlinkName = l
		} else if strings.HasPrefix(l, "@") {
			authorName = l
			permlinkCounter = i
		}
	}

	if authorName != "" && permlinkName != "" {
		url := fmt.Sprintf("https://api.steemhunt.com/posts/%s/%s.json", authorName, permlinkName)
		resp, err := http.CG.MakeReq(url)
		if err != nil {
			em := d.GenErrorMessage("Post not found or SteemHunt API down.")
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}
		var data SteemHuntPost
		err = json.Unmarshal(resp, &data)
		if err != nil {
			em := d.GenErrorMessage("Unable to process post data from SteemHunt API.")
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}

		sender, _ := d.Discord.UserChannelCreate(m.Author.ID)
		// Verify?
		var verifyMessage string
		if data.IsVerified {
			verifyMessage = fmt.Sprintf(":white_check_mark: by %s", data.VerifiedBy)
		} else {
			verifyMessage = ":x:"
		}
		// Timeago
		postTime, _ := timeago.TimeAgoFromNowWithTime(data.CreatedAt)

		em := d.GenMultipleEmbed(
			d.Green,
			data.Title,
			[]*discord.MessageEmbedField{
				&discord.MessageEmbedField{
					Name:   "Author",
					Value:  data.Author,
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Tagline",
					Value:  data.Tagline,
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "URL",
					Value:  fmt.Sprintf("https://steemhunt.com/%s/%s?ref=superoo7", authorName, permlinkName),
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Payout",
					Value:  fmt.Sprintf("$%f", data.PayoutValue),
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Vote",
					Value:  fmt.Sprintf("%d 👍", len(data.ActiveVotes)),
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Hunt Score",
					Value:  fmt.Sprintf("%d :regional_indicator_s: :regional_indicator_h:", len(data.ActiveVotes)),
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Verify?",
					Value:  verifyMessage,
					Inline: true,
				},
				&discord.MessageEmbedField{
					Name:   "Time",
					Value:  fmt.Sprintf("%s (%s KST)", postTime, data.CreatedAt.Format(time.Kitchen)),
					Inline: true,
				},
			},
		)
		c <- d.DiscordEmbedMessage{CID: sender.ID, Message: em}
	}

}
