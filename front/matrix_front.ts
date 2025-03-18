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

  const detButton: HTMLInputElement = document.createElement("input");
  detButton.value = "Find determinant";
  detButton.type = "button";
  divElement.appendChild(detButton);
  detButton.addEventListener("click", getMatrix);
}

function getMatrix() {
  let matrix: number[][] = [];
  for (let i: number = 0; i < size; i++) {
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

function generateLinearSystem() {
  const sizeInput = document.getElementById("sizeInput") as HTMLInputElement;
  const divElement = document.getElementById(
    "matrixContainer"
  ) as HTMLDivElement;
  size = parseInt(sizeInput.value);

  divElement.innerHTML = "";

  for (let i: number = 0; i < size; i++) {
    const row = document.createElement("div");
    for (let j = 0; j <= size; j++) {
      const input: HTMLInputElement = document.createElement("input");
      input.type = "text";
      if (j != size) {
        input.id = `${i}${j}`;
      } else {
        input.id = `b${i}`;
        const addEqual: HTMLSpanElement = document.createElement("span");
        addEqual.textContent = "=";
        row.appendChild(addEqual);
      }
      row.appendChild(input);
    }
    divElement.appendChild(row);
  }

  const linearSystemButton: HTMLInputElement = document.createElement("input");
  linearSystemButton.type = "button";
  linearSystemButton.value = "Solve system";
  linearSystemButton.addEventListener("click", getLinearSystem);
  divElement.appendChild(linearSystemButton);
}

function getLinearSystem() {
  let matrix: number[][] = [];
  let b: number[] = [];

  for (let i: number = 0; i < size; i++) {
    matrix[i] = [];

    for (let j: number = 0; j <= size; j++) {
      if (j != size) {
        const elmOfMatrix = document.getElementById(
          `${i}${j}`
        ) as HTMLInputElement;
        matrix[i][j] = parseFloat(elmOfMatrix.value);
      } else {
        const elmOfB = document.getElementById(`b${i}`) as HTMLInputElement;
        b[i] = parseFloat(elmOfB.value);
      }
    }
  }

  console.log(matrix);
  console.log(b);

  fetch("http://localhost:8080/linearsystem", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      matrix: matrix,
      b: b,
    }),
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
      console.log("Erro ao enviar sistema linear", error);
    });
}
