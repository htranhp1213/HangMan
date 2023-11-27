
// Initial constant values to get the needed elements
const letterContainer = document.getElementById("letter-container");
const optionsContainer = document.getElementById("options-container");
const userInputSection = document.getElementById("user-input-section");
const newGameContainer = document.getElementById("new-game-container");
const newGameButton = document.getElementById("new-game-button");
const canvas = document.getElementById("canvas");
const resultText = document.getElementById("result-text");

// All word options for the specific categories
let options = {

  fruits: [
    "Apple",
    "Blueberry",
    "Mandarin",
    "Pineapple",
    "Pomegranate",
    "Watermelon",
  ],
  animals: ["Hedgehog", "Rhinoceros", "Squirrel", "Panther", "Walrus", "Zebra"],
  countries: [
    "India",
    "Hungary",
    "Kyrgyzstan",
    "Switzerland",
    "Zimbabwe",
    "Dominica",
  ],

};

// Win counter, game counter and current chosen word
let winCount = 0;
let count = 0;
let chosenWord = "";

// Display option buttons
const displayOptions = () => {

  optionsContainer.innerHTML += `<h3>Please Select An Option</h3>`;
  let buttonCon = document.createElement("div");

  for (let value in options) {

    buttonCon.innerHTML += `<button class="options" onclick="generateWord('${value}')">${value}</button>`;

  }

  optionsContainer.appendChild(buttonCon);

};

// Disables all buttons so user can no longer press them
const blocker = () => {

  let optionsButtons = document.querySelectorAll(".options");
  let letterButtons = document.querySelectorAll(".letters");

  optionsButtons.forEach((button) => {

    button.disabled = true;

  });

  letterButtons.forEach((button) => {

    button.disabled.true;

  });

  newGameContainer.classList.remove("hide");

};

// Generates a random word from the slice based on the chosen category
const generateWord = (optionValue) => {

  let optionsButtons = document.querySelectorAll(".options");

  optionsButtons.forEach((button) => {

    if (button.innerText.toLowerCase() === optionValue) {

      button.classList.add("active");

    }

    button.disabled = true;

  });

  // Resets the UI
  letterContainer.classList.remove("hide");
  userInputSection.innerText = "";

  // Fetch random word from our server
  let optionArray = options[optionValue];
  chosenWord = optionArray[Math.floor(Math.random() * optionArray.length)];
  chosenWord = chosenWord.toUpperCase();

  fetch(`http://localhost:8080/generate-word?option=${optionValue}`)

    .then(response => response.json())
    .then(data => {

      chosenWord = data.chosenWord;

      // Console logs for debugging and presentation purposes
      console.log('Chosen Word:', chosenWord);
      console.log('Word Length:', chosenWord.length);
      console.log('Individual Characters:', Array.from(chosenWord));

      // Displays dashes for each letter of the word
      let displayItem = '';
      for (let i = 0; i < chosenWord.length; i++) {

        displayItem += '<span class="dashes">_</span>';

      }

      userInputSection.innerHTML = displayItem;
      
    })

  .catch(error => console.error('Error fetching word:', error));

};

// Initial Function (Called when page loads/user presses new game)
const initializer = () => {

  winCount = 0;
  count = 0;

  // Resets the UI elements
  userInputSection.innerHTML = "";
  optionsContainer.innerHTML = "";
  letterContainer.classList.add("hide");
  newGameContainer.classList.add("hide");
  letterContainer.innerHTML = "";

  // Creates individual buttons for letters A-Z
  for (let i = 65; i < 91; i++) {

    let button = document.createElement("button");
    button.classList.add("letters");
    button.innerText = String.fromCharCode(i);
    button.addEventListener("click", () => {

      let charArray = chosenWord.split("");
      let dashes = document.getElementsByClassName("dashes");

      // If array contains clciked value replace the matched dash with letter else dram on canvas
      if (charArray.includes(button.innerText)) {

        charArray.forEach((char, index) => {

          // If character in array is same as clicked button
          if (char === button.innerText) {

            // Replace dash with letter
            dashes[index].innerText = char;

            // Increment counter
            winCount += 1;

            // If winCount equals word length
            if (winCount == charArray.length) {

              // Display win message
              resultText.innerHTML = `<h2 class='win-msg'>You Win!!</h2><p>The word was <span>${chosenWord}</span></p>`;
              
              // Block all buttons
              blocker();

            }

          }

        });

      } else {

        // Increment game count and draw corresponding hangman piece
        count += 1;
        drawMan(count);

        // Count==6 because head,body,left arm, right arm,left leg,right leg
        if (count == 6) {

          // Display lose message
          resultText.innerHTML = `<h2 class='lose-msg'>You Lose!!</h2><p>The word was <span>${chosenWord}</span></p>`;

          // Block all buttons
          blocker();

        }

      }

      // Disable clicked button
      button.disabled = true;

    });

    letterContainer.append(button);

  }

  // Display categories
  displayOptions();

  // Call to canvasCreator (for clearing previous canvas and creating initial canvas)
  let { initialDrawing } = canvasCreator();

  initialDrawing();

};

// Function to create canvas and manage drawing on said canvas
const canvasCreator = () => {

  let context = canvas.getContext("2d");
  context.beginPath();
  context.strokeStyle = "#000";
  context.lineWidth = 2;

  // For drawing lines
  const drawLine = (fromX, fromY, toX, toY) => {

    context.moveTo(fromX, fromY);
    context.lineTo(toX, toY);
    context.stroke();

  };

  const head = () => {

    context.beginPath();
    context.arc(70, 30, 10, 0, Math.PI * 2, true);
    context.stroke();

  };

  const body = () => {

    drawLine(70, 40, 70, 80);

  };

  const leftArm = () => {

    drawLine(70, 50, 50, 70);

  };

  const rightArm = () => {

    drawLine(70, 50, 90, 70);

  };

  const leftLeg = () => {

    drawLine(70, 80, 50, 110);

  };

  const rightLeg = () => {

    drawLine(70, 80, 90, 110);

  };

  // Initial frame drawing
  const initialDrawing = () => {

    // Clear canvas
    context.clearRect(0, 0, context.canvas.width, context.canvas.height);

    // Bottom line
    drawLine(10, 130, 130, 130);

    // Left line
    drawLine(10, 10, 10, 131);

    // Top line
    drawLine(10, 10, 70, 10);

    // Small top line
    drawLine(70, 10, 70, 20);

  };

  return { initialDrawing, head, body, leftArm, rightArm, leftLeg, rightLeg };

};

// Draws the hangman
const drawMan = (count) => {

  let { head, body, leftArm, rightArm, leftLeg, rightLeg } = canvasCreator();
  switch (count) {

    case 1:
      head();
      break;

    case 2:
      body();
      break;

    case 3:
      leftArm();
      break;

    case 4:
      rightArm();
      break;

    case 5:
      leftLeg();
      break;

    case 6:
      rightLeg();
      break;

    default:
      break;

  }

};

// New Game
newGameButton.addEventListener("click", initializer);
window.onload = initializer;
