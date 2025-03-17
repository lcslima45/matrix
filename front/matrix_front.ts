let size: number;
function generateMatrix() {
  const sizeInput = document.getElementById("sizeInput") as HTMLInputElement;
  const divElement = document.getElementById(
    "matrixContainer"
  ) as HTMLDivElement;
  size = parseInt(sizeInput.value);

  divElement.innerHTML = "";

  for (let i = 0; i < size; i++) {
    const row = document.createElement("div");
    for (let j = 0; j < size; j++) {
      const input = document.createElement("input");
      input.id = `${i}${j}`;
      input.type = "text";
      row.appendChild(input);
    }
    divElement.appendChild(row);
  }

  const detButton = document.createElement("input");
  detButton.value = "Find determinant";
  detButton.type = "button";
  divElement.appendChild(detButton);
  detButton.addEventListener("click", getMatrix);
}

function getMatrix() {
  let matrix: number[][] = [];
  for (let i = 0; i < size; i++) {
    matrix[i] = [];
    for (let j = 0; j < size; j++) {
      const elementOfMatrix = document.getElementById(
        `${i}${j}`
      ) as HTMLInputElement;
      matrix[i][j] = parseFloat(elementOfMatrix.value);
    }
  }

  fetch("http://localhost:8080/determinant", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ matrix: matrix }),
  })
    .then((response) => response.json())
    .then((data) => {
      const responseContainer = document.createElement("div");
      responseContainer.innerHTML = `<div> Resposta </br> ${JSON.stringify(
        data
      )} </br> </div>`;
      const matrixContainer = document.getElementById(
        "matrixContainer"
      ) as HTMLDivElement;
      matrixContainer.appendChild(responseContainer);
    })
    .catch((error) => {
      console.log("Erro ao enviar a matriz", error);
    });
}
