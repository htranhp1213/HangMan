const pumpkinHeadButton = document.getElementById("pumpkin-head-button");
const head = document.getElementById("head");

pumpkinHeadButton.addEventListener("click", changeHeadToPumpkin);

function changeHeadToPumpkin() {
    head.style.backgroundColor = "#FF6347"; // Change to pumpkin color
    pumpkinHeadButton.disabled = true; // Disable the button after changing the head
}
