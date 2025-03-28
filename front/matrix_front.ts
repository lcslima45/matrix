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
      if (j !== size) {
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

function generateMatrixSum() {
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
      input.id = `a${i}${j}`;
      input.type = "text";
      row.appendChild(input);
    }
    divElement.appendChild(row);
    if (i === size-1) {
      const sumEqual: HTMLSpanElement = document.createElement("span");
      sumEqual.textContent = "+";
      row.appendChild(sumEqual);
    }
  }
  
  for (let i = 0; i < size; i++) {
    const row = document.createElement("div");
    for (let j = 0; j < size; j++) {
      const input = document.createElement("input");
      input.id = `b${i}${j}`;
      input.type = "number"; // Alterado para 'number' para garantir que seja um número
      row.appendChild(input);
    }
    divElement.appendChild(row);
  }

  const detButton: HTMLInputElement = document.createElement("input");
  detButton.value = "Sum Matrix";
  detButton.type = "button";
  divElement.appendChild(detButton);
  detButton.addEventListener("click", getMatrixSum);
}


function getMatrixSum() {
  let matrixA: number[][] = [];
  let matrixB: number[][] = [];
  
  for (let i: number = 0; i < size; i++) {
    matrixA[i] = [];
    matrixB[i] = [];
    for (let j = 0; j < size; j++) {
      const elementOfMatrixA = document.getElementById(
        `a${i}${j}`
      ) as HTMLInputElement;
      const elementOfMatrixB = document.getElementById(
        `b${i}${j}`
      ) as HTMLInputElement;
      matrixA[i][j] = parseFloat(elementOfMatrixA.value);
      matrixB[i][j] = parseFloat(elementOfMatrixB.value);
    }
  }
  fetch("http://localhost:8080/sum", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ matrixA: matrixA, matrixB: matrixB }),
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
      console.log("Erro ao enviar a soma", error);
    });
}

function generateMatrixProduct() {
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
      input.id = `a${i}${j}`;
      input.type = "text";
      row.appendChild(input);
    }
    divElement.appendChild(row);
    if (i === size-1) {
      const sumEqual: HTMLSpanElement = document.createElement("span");
      sumEqual.textContent = "X";
      row.appendChild(sumEqual);
    }
  }
  
  for (let i = 0; i < size; i++) {
    const row = document.createElement("div");
    for (let j = 0; j < size; j++) {
      const input = document.createElement("input");
      input.id = `b${i}${j}`;
      input.type = "number"; // Alterado para 'number' para garantir que seja um número
      row.appendChild(input);
    }
    divElement.appendChild(row);
  }

  const detButton: HTMLInputElement = document.createElement("input");
  detButton.value = "Make Product";
  detButton.type = "button";
  divElement.appendChild(detButton);
  detButton.addEventListener("click", getMatrixProduct);
}


function getMatrixProduct() {
  let matrixA: number[][] = [];
  let matrixB: number[][] = [];
  
  for (let i: number = 0; i < size; i++) {
    matrixA[i] = [];
    matrixB[i] = [];
    for (let j = 0; j < size; j++) {
      const elementOfMatrixA = document.getElementById(
        `a${i}${j}`
      ) as HTMLInputElement;
      const elementOfMatrixB = document.getElementById(
        `b${i}${j}`
      ) as HTMLInputElement;
      matrixA[i][j] = parseFloat(elementOfMatrixA.value);
      matrixB[i][j] = parseFloat(elementOfMatrixB.value);
    }
  }
  fetch("http://localhost:8080/product", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ matrixA: matrixA, matrixB: matrixB }),
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
      console.log("Erro ao enviar o produto", error);
    });
}

function generateMatrixLU() {
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
      input.id = `a${i}${j}`;
      input.type = "text";
      row.appendChild(input);
    }
    divElement.appendChild(row);
  }
  const detButton: HTMLInputElement = document.createElement("input");
  detButton.value = "Make LU Decomposition";
  detButton.type = "button";
  divElement.appendChild(detButton);
  detButton.addEventListener("click", getMatrixLU);
}


function getMatrixLU() {
  let matrix: number[][] = [];
  
  for (let i: number = 0; i < size; i++) {
    matrix[i] = [];
    for (let j = 0; j < size; j++) {
      const elementOfMatrixA = document.getElementById(
        `a${i}${j}`
      ) as HTMLInputElement;
      matrix[i][j] = parseFloat(elementOfMatrixA.value);
    }
  }
  console.log(matrix)
  fetch("http://localhost:8080/ludecompose", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ matrix: matrix}),
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
      console.log("Erro ao enviar o LU", error);
    });
}