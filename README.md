# Word Scramble

Welcome to **Word Scramble**, a simple yet fun game where you test your word-guessing skills! The game randomly selects a word, scrambles its letters, and challenges you to guess the correct word within three attempts.

## Features

- Play directly in the command-line interface (CLI).
- Option to run a web server to play via a web browser.
- Configure the web server port using command-line arguments.

## How to Play

### CLI Mode (Default)

Simply run the game without any flags to start playing in the terminal:

```bash
$ ./word-scramble
```

The game will display a scrambled word, and you'll have three chances to guess the correct word.

### Web Mode

To start the web server, use the `-web` flag:

```bash
$ ./word-scramble -web
```

#### Custom Port

By default, the server runs on port `8080`. You can specify a custom port using the `-port` flag:

```bash
$ ./word-scramble -web -port 9090
```

Then, visit [http://localhost:9090](http://localhost:9090) to play.

## Command-Line Arguments

- `-web`: Starts the web server for the game.
- `-port <number>`: Specifies the port for the web server.

## Example Game Flow

### CLI

1. The game displays a scrambled word: `elppa`
2. You guess: `apple`
3. If correct, the game congratulates you. If incorrect, you have up to 3 attempts.

### Web

1. Open the web page.
2. Enter your guesses in the provided input box.
3. Get instant feedback on your guesses.

## Installation

Make sure you have Go installed on your system. Build the game with:

```bash
go build -o word-scramble
```

Alternatively, you can download the pre-built packages from the [GitHub Releases](https://github.com/MementoMori11723/Word-Scramble/releases/tag/v0.0.1) page.

## Development

If you want to contribute or modify the code:

1. Clone the repository.
2. Make changes and test them.
3. Submit a pull request if you're contributing.

## License

This project is open-source and available under the [MIT License](LICENSE).

Enjoy the game and happy scrambling!
