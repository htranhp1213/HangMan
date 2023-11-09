const attemptsLeftSpan = document.getElementById("attemptsLeft");
const wordDisplayDiv = document.getElementById("wordDisplay");
const guessInput = document.getElementById("guessInput");
const messageParagraph = document.getElementById("message");
const pumpkinHeadButton = document.getElementById("pumpkin-head-button");
const head = document.getElementById("head");
const makeGuessButton = document.getElementById("makeGuessButton");

pumpkinHeadButton.addEventListener("click", changeHeadToPumpkin);

function changeHeadToPumpkin() {
    head.style.backgroundColor = "#FF6347"; // Change to pumpkin color
    pumpkinHeadButton.disabled = true; // Disable the button after changing the head
}
function updateAttemptsLeft(attemptsLeft) {
    attemptsLeftSpan.textContent = attemptsLeft;
}

function updateWordDisplay(word) {
    wordDisplayDiv.textContent = word;
}

function showMessage(message) {
    messageParagraph.textContent = message;
}
