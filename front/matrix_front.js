var size;
function generateMatrix() {
    var sizeInput = document.getElementById("sizeInput");
    var divElement = document.getElementById("matrixContainer");
    size = parseInt(sizeInput.value);
    divElement.innerHTML = "";
    for (var i = 0; i < size; i++) {
        var row = document.createElement("div");
        for (var j = 0; j < size; j++) {
            var input = document.createElement("input");
            input.id = "".concat(i).concat(j);
            input.type = "text";
            row.appendChild(input);
        }
        divElement.appendChild(row);
    }
    var detButton = document.createElement("input");
    detButton.value = "Find determinant";
    detButton.type = "button";
    divElement.appendChild(detButton);
    detButton.addEventListener("click", getMatrix);
}
function getMatrix() {
    var matrix = [];
    for (var i = 0; i < size; i++) {
        matrix[i] = [];
        for (var j = 0; j < size; j++) {
            var elementOfMatrix = document.getElementById("".concat(i).concat(j));
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
        .then(function (response) { return response.json(); })
        .then(function (data) {
        var responseContainer = document.createElement("div");
        responseContainer.innerHTML = "<div> Resposta </br> ".concat(JSON.stringify(data), " </br> </div>");
        var matrixContainer = document.getElementById("matrixContainer");
        matrixContainer.appendChild(responseContainer);
    })
        .catch(function (error) {
        console.log("Erro ao enviar a matriz", error);
    });
}
function generateLinearSystem() {
    var sizeInput = document.getElementById("sizeInput");
    var divElement = document.getElementById("matrixContainer");
    size = parseInt(sizeInput.value);
    divElement.innerHTML = "";
    for (var i = 0; i < size; i++) {
        var row = document.createElement("div");
        for (var j = 0; j <= size; j++) {
            var input = document.createElement("input");
            input.type = "text";
            if (j != size) {
                input.id = "".concat(i).concat(j);
            }
            else {
                input.id = "b".concat(i);
                var addEqual = document.createElement("span");
                addEqual.textContent = "=";
                row.appendChild(addEqual);
            }
            row.appendChild(input);
        }
        divElement.appendChild(row);
    }
    var linearSystemButton = document.createElement("input");
    linearSystemButton.type = "button";
    linearSystemButton.value = "Solve system";
    linearSystemButton.addEventListener("click", getLinearSystem);
    divElement.appendChild(linearSystemButton);
}
function getLinearSystem() {
    var matrix = [];
    var b = [];
    for (var i = 0; i < size; i++) {
        matrix[i] = [];
        for (var j = 0; j <= size; j++) {
            if (j != size) {
                var elmOfMatrix = document.getElementById("".concat(i).concat(j));
                matrix[i][j] = parseFloat(elmOfMatrix.value);
            }
            else {
                var elmOfB = document.getElementById("b".concat(i));
                b[i] = parseFloat(elmOfB.value);
            }
        }
    }
    console.log(matrix);
    console.log(b);
    fetch("http://localhost:8080/determinant", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            matrix: matrix,
            b: b,
        }),
    })
        .then(function (response) { return response.json(); })
        .then(function (data) {
        var responseContainer = document.createElement("div");
        responseContainer.innerHTML = "<div> Resposta </br> ".concat(JSON.stringify(data), " </br> </div>");
        var matrixContainer = document.getElementById("matrixContainer");
        matrixContainer.appendChild(responseContainer);
    })
        .catch(function (error) {
        console.log("Erro ao enviar sistema linear", error);
    });
}
