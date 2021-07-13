package osu

import (
	"Digobo/config"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

type EventType string

const (
	/*
	 * Provides: Event.Achievement, Event.User
	 */
	EVENT_TYPE_ACHIEVEMENT EventType = "achievement"

	/*
	 * Provides: Event.Beatmap, Event.Count
	 */
	EVENT_TYPE_BEATMAP_PLAYCOUNT EventType = "beatmapPlaycount"

	/*
	 * Provides: Event.Approval, Event.Beatmapset, Event.User
	 */
	EVENT_TYPE_BEATMAPSET_APPROVE EventType = "beatmapsetApprove"

	/*
	 * Provides: Event.Beatmapset
	 */
	EVENT_TYPE_BEATMAPSET_DELETE EventType = "beatmapsetDelete"

	/*
	 * Provides: Event.Beatmapset, Event.User
	 */
	EVENT_TYPE_BEATMAPSET_REVIVE EventType = "beatmapsetRevive"

	/*
	 * Provides: Event.Beatmapset, Event.User
	 */
	EVENT_TYPE_BEATMAPSET_UPDATE EventType = "beatmapsetUpdate"

	/*
	 * Provides: Event.Beatmapset, Event.User
	 */
	EVENT_TYPE_BEATMAPSET_UPLOAD EventType = "beatmapsetUpload"

	/*
	 * Provides: Event.ScoreRank, Event.Rank, Event.Mode, Event.Beatmap, Event.User
	 */
	EVENT_TYPE_RANK EventType = "rank"

	/*
	 * Provides: Event.Mode, Event.User, Event.Beatmap
	 */
	EVENT_TYPE_RANK_LOST EventType = "rankLost"

	/*
	 * Provides: Event.User
	 */
	EVENT_TYPE_USER_SUPPORT_AGAIN EventType = "userSupportAgain"

	/*
	 * Provides: Event.User
	 */
	EVENT_TYPE_USER_SUPPORT_FIRST EventType = "userSupportFirst"

	/*
	 * Provides: Event.User
	 */
	EVENT_TYPE_USER_SUPPORT_GIFT EventType = "userSupportGift"

	/*
	 * Provides: Event.User
	 */
	EVENT_TYPE_USERNAME_CHANGE EventType = "usernameChange"
)

type Approval string

const (
	APPROVAL_RANKED    Approval = "ranked"
	APPROVAL_APPROVED  Approval = "approved"
	APPROVAL_QUALIFIED Approval = "qualified"
	APPROVAL_LOVED     Approval = "loved"
)

type Event struct {
	CreatedAt   time.Time   `json:"created_at"`
	Id          int         `json:"id"`
	Type        EventType   `json:"type"`
	CreatedAt1  time.Time   `json:"createdAt,omitempty"`
	ScoreRank   string      `json:"scoreRank,omitempty"`
	Rank        int         `json:"rank,omitempty"`
	Mode        string      `json:"mode,omitempty"`
	Count       int         `json:"count,omitempty"`
	Approval    Approval    `json:"approval,omitempty"`
	Beatmap     Beatmap     `json:"beatmap,omitempty"`
	Beatmapset  Beatmapset  `json:"beatmapset,omitempty"`
	User        User        `json:"user,omitempty"`
	Achievement Achievement `json:"achievement,omitempty"`
}

func (this *Event) PrepareEmbed(userName string, userId int) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title:       userName,
		Description: this.createTypeTitle(),
		Timestamp:   this.CreatedAt.Format(time.RFC3339),
		Color:       config.Config.Bot.DefaultEmbedColor,
		Author:      &discordgo.MessageEmbedAuthor{},
		Fields:      this.createMessageEmbedFields(),
	}

	embed.URL = "https://osu.ppy.sh/users/" + strconv.Itoa(userId)

	return embed
}

func (this *Event) createTypeTitle() string {
	switch this.Type {
	case EVENT_TYPE_ACHIEVEMENT:
		return "New Achievement"
	case EVENT_TYPE_BEATMAP_PLAYCOUNT:
		return "Reached a high playcount"
	case EVENT_TYPE_BEATMAPSET_APPROVE:
		return "Got a new beatmap set approved"
	case EVENT_TYPE_BEATMAPSET_DELETE:
		return "Deleted a beatmap set"
	case EVENT_TYPE_BEATMAPSET_REVIVE:
		return "Revived a beatmap set"
	case EVENT_TYPE_BEATMAPSET_UPDATE:
		return "Updated a beatmap set"
	case EVENT_TYPE_BEATMAPSET_UPLOAD:
		return "Uploaded a new beatmap set"
	case EVENT_TYPE_RANK:
		return "Scored a new rank"
	case EVENT_TYPE_RANK_LOST:
		return "Lost a rank"
	case EVENT_TYPE_USER_SUPPORT_AGAIN:
		return "Supports osu! again"
	case EVENT_TYPE_USER_SUPPORT_FIRST:
		return "Supports osu! for the first time"
	case EVENT_TYPE_USER_SUPPORT_GIFT:
		return "Got osu! support gifted"
	case EVENT_TYPE_USERNAME_CHANGE:
		return "Changed their username"
	default:
		return "Unknown type"
	}
}

func (this *Event) getApprovalState() string {
	switch this.Approval {
	case APPROVAL_RANKED:
		return "ranked"
	case APPROVAL_APPROVED:
		return "approved"
	case APPROVAL_QUALIFIED:
		return "qualified"
	case APPROVAL_LOVED:
		return "loved"
	default:
		return "unknown approval state"
	}
}

func (this *Event) createMessageEmbedFields() []*discordgo.MessageEmbedField {
	var ret []*discordgo.MessageEmbedField

	switch this.Type {
	case EVENT_TYPE_ACHIEVEMENT:
		prepend := ""
		if this.Achievement.Instructions != "" {
			prepend = "\n (" + this.Achievement.Instructions + ")"
		}

		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Achievement.Name,
			Value:  this.Achievement.Description + prepend,
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAP_PLAYCOUNT:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Beatmap.Title + " played for " + strconv.Itoa(this.Count) + " times",
			Value:  OSU_API_URL + this.Beatmap.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAPSET_APPROVE:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Beatmapset.Title + " new status: " + this.getApprovalState(),
			Value:  OSU_API_URL + this.Beatmapset.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAPSET_DELETE:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Beatmapset.Title + " got deleted",
			Value:  "We are very sad :(",
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAPSET_REVIVE:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Beatmapset.Title + " got revived",
			Value:  OSU_API_URL + this.Beatmapset.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAPSET_UPDATE:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   this.Beatmapset.Title + " got an update",
			Value:  OSU_API_URL + this.Beatmapset.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_BEATMAPSET_UPLOAD:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "New map: " + this.Beatmapset.Title,
			Value:  OSU_API_URL + this.Beatmapset.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_RANK:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "scored a new rank (Rank: #" + strconv.Itoa(this.Rank) + ", ScoreRank: " + this.ScoreRank + ")",
			Value:  OSU_API_URL + this.Beatmap.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_RANK_LOST:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "Lost his #1 rank :(",
			Value:  OSU_API_URL + this.Beatmap.Url,
			Inline: false,
		})
		break
	case EVENT_TYPE_USER_SUPPORT_AGAIN:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "Supports osu! again",
			Value:  "we like that!",
			Inline: false,
		})
		break
	case EVENT_TYPE_USER_SUPPORT_FIRST:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "Supports osu! for the first time",
			Value:  "great Achievement!",
			Inline: false,
		})
		break
	case EVENT_TYPE_USER_SUPPORT_GIFT:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "Got osu! supporter gifted",
			Value:  "Go ahead and say thank you!",
			Inline: false,
		})
		break
	case EVENT_TYPE_USERNAME_CHANGE:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   "Changed his username",
			Value:  "Username changed from: " + this.User.PreviousUsername + " to: " + this.User.Username,
			Inline: false,
		})
		break
	default:
		ret = append(ret, &discordgo.MessageEmbedField{
			Name:   string(this.Type) + " is currently not known",
			Value:  "Go ahead and implement it you lazy...",
			Inline: false,
		})
		break
	}

	return ret
}
