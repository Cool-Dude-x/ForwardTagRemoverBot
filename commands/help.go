


package commands

import (
 "fmt"

 "github.com/PaulSonOfLars/gotgbot"
 "github.com/PaulSonOfLars/gotgbot/ext"
 "github.com/PaulSonOfLars/gotgbot/parsemode"
 "go.uber.org/zap"
)

func Start(b ext.Bot, u *gotgbot.Update) error {

 startButton := [][]ext.InlineKeyboardButton{make([]ext.InlineKeyboardButton, 2), make([]ext.InlineKeyboardButton, 1)}

 startButton[0][0] = ext.InlineKeyboardButton{
  Text: "Join Group 1 ",
  Url:  "https://t.me/missioncrackjeeneet",
 }

 startButton[0][1] = ext.InlineKeyboardButton{
  Text: "Join Group 2 ",
  Url:  "https://t.me/studymaterialmcjn",
 }

 startButton[1][0] = ext.InlineKeyboardButton{
  Text: "My Owner",
  Url:  "https://t.me/tausifur_123",
 }

 markup := ext.InlineKeyboardMarkup{InlineKeyboard: &startButton}

 msg := b.NewSendableMessage(u.EffectiveChat.Id, fmt.Sprintf("Hi [%s](tg://user?id=%v)\nI am A Forward Tag remover Bot.Send /help To Know What I Can Do", u.EffectiveUser.FirstName, u.EffectiveUser.Id))
 msg.ReplyToMessageId = u.EffectiveMessage.MessageId
 msg.ReplyMarkup = &markup
 msg.ParseMode = parsemode.Markdown
 _, err := msg.Send()
 if err != nil {
  b.Logger.Warnw("Error in sending", zap.Error(err))
 }
 return err
}
