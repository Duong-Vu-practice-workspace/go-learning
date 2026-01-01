[Back](https://www.boot.dev/lessons/5498d6da-38ea-4993-bc88-cc5f07eff909)

[Next](https://www.boot.dev/lessons/7af7c74b-99bb-4420-9fa4-275fb8ec7a5a)

## Message Formatter

As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user's preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.

## Assignment

1. Implement the `formatter` interface with a method `format` that returns a formatted string.
2. Define structs that satisfy the `formatter` interface: `plainText`, `bold`, `code`.
    - The structs must all have a `message` field of type `string`
- `plainText` should return the message as is.
- `bold` should wrap the message in two asterisks (\*\*) to simulate bold text (e.g., `**message**`).
- `code` should wrap the message in a single backtick (\`) to simulate code block (e.g., `` `message` ``)

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Bane of End-Users, can assist... *for a price*.

<audio xmlns="http://www.w3.org/1999/xhtml"></audio>

package main



// Don't Touch below this line



func sendMessage (format formatter) string {

return format.format () // Adjusted to call Format without an argument

}