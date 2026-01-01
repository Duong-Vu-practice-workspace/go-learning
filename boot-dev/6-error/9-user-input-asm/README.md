[Back](https://www.boot.dev/lessons/c6c9a896-4816-48f2-ae5c-16b3a117cd7f)

[Next](https://www.boot.dev/lessons/d8b6aaab-5a7c-4fb9-b8d8-82297029057a)

## User Input

In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.

## Assignment

Complete the `validateStatus` function. It should return an error when the status update violates any of the rules:

- If the status is empty, return an error that reads `status cannot be empty`.
- If the status exceeds 140 characters, return an error that says `status exceeds 140 characters`.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Sleepy Spellcaster, can assist... *for a price*.

<audio xmlns="http://www.w3.org/1999/xhtml"></audio>

package main



import (

"errors"

)



func validateStatus (status string) error {

//?

}