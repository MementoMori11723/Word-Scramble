const [word, arr] = ["Hello World", []],
  readline = require("readline").createInterface({
    input: process.stdin,
    output: process.stdout,
  });
word.split("").map((e) => arr.push(e));
const shuffle = (e) => {
  for (let r = e.length - 1; r > 0; r--) {
    let o = Math.floor(Math.random() * (r + 1));
    [e[r], e[o]] = [e[o], e[r]];
  }
  return e;
};
(newarr = shuffle(arr)),
  console.log("Guess the word : \nStart = '" + newarr + "'"),
  readline.question("What's your Guess : ", (e) => {
    e === word ? console.log("You Win!") : console.log("You Lose!"),
      readline.close();
  });
