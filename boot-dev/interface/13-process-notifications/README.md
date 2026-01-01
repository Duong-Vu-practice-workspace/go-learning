[Back](https://www.boot.dev/lessons/4988b18c-4da4-4f55-84fb-25096dcc66c2)

[Next](https://www.boot.dev/lessons/1aac4f6b-8416-4a67-a1f7-fb46bcb11565)

## Process Notifications

Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.

## Assignment

1. 1.
2.
3.
2. 1.
2.
3.
4.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Primeval 10x Developer, can assist... *for a price*.

<audio xmlns="http://www.w3.org/1999/xhtml"></audio>

package main



type notification interface {

importance () int

}



type directMessage struct {

senderUsername string

messageContent string

priorityLevel int

isUrgent bool

}



type groupMessage struct {

groupName string

messageContent string

priorityLevel int

}



type systemAlert struct {

alertCode string

messageContent string

}



//?



func processNotification (n notification) (string,int) {

//?

}