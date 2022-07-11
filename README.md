<h1 align="center">Grunt - Discord Bot</h1>

<p align="center">
    <img src="https://raw.githubusercontent.com/jostgrant/grunt/main/asset/grunt.gif" " alt="Grunt's Logo / Mascot.">
</p>

<h3 align="center"> You can invite the bot to your server using <a href="bot-invite-link">this</a> link<h3>

<h5 align="center"> Want to contribute to the bot in some way? Read the <a href="CONTRIB.md">CONTRIB.md</a> file<h5>

---

### Table of contents
- [The Pitch](#the-pitch)
- [Commands](#commands)
- [Run It Yourself](#run-it-yourself)
- [Credit And Thanks](#credit-and-thanks)
--- 

### The Pitch
**DESCRIPTION OF YOUR BOT**: 
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque sit amet consequat magna. Nam malesuada vitae tortor id aliquam. Praesent congue lectus quis neque scelerisque euismod. Pellentesque malesuada diam ut purus facilisis, et ultricies tellus vulputate. Sed a consequat ipsum. Vivamus dui neque, venenatis non dui vitae, mattis pulvinar nisl. Praesent vitae viverra mi.

---

### Commands
All the commands start with the prefix `/`

| Name              | Description                               |
|-------------------|-------------------------------------------|
| hello             | Say hello to the bot!                     |
| ...               | ...                                       |
| ...               | ...                                       |

---

### Run It Yourself
If you want to compile the bot's code to then run it in a local server, you will need:

| Name     | Version | Use | Links|
|----------|---------|-----|------|
| Git      | v0.9+   | Clonning the repository to your computer | [install](https://git-scm.com/downloads) [docs](https://git-scm.com/doc)  |
| Go tool  | v0.5.0+ | Building the actual bot as an executable | [install](https://golang.org) [docs](https://golang.org/cmd/go/)          |

Then, run open a terminal / cmd and run:
###### Clone the repository
```shell
git clone https://github.com/jostgrant/grunt.git
```

###### Move inside the repository
```shell
cd grunt/
```
###### Use the Golang tool to build the code
```shell
go build -o bot
```
###### Execute the code
```shell
./bot -t <bot token>
```

---

### Credit And Thanks
Special thanks to [Lieuwe Berg](https://github.com/lieuweberg) & [Pablo Corbalán](https://github.com/pablocorbalann) for their DiscordGo Templates; Where the code and beginner friendly readme for this project was seeded from (respectively). 
And of course thanks to [Bruce Marriner](https://github.com/bwmarrin/discordgo) and [everyone else](https://github.com/bwmarrin/discordgo/graphs/contributors) who's made DiscordGo possible!
