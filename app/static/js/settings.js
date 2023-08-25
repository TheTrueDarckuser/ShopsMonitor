function XHRrequest(method, url, data) {
    return new Promise(function(resolve) {
        const xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function(e) {
          if (xhr.readyState === 4 && xhr.status == 200) {
              resolve(xhr.response);
          }
        }
        xhr.open(method, url, true);
        xhr.send(data);
      })
}

function AJAXrequest(method, url, data, f) {
    $.ajax({
        type: method,
        url: url,
        data: data,
        success: function(response) {
            f(response);
        }
    });
}

// -------------------------------- Setup -------------------------------- //

window.onload = async function() {
    await AJAXrequest("GET", "/api/patterns/", null, applyPatterns);
}

async function applyPatterns(response) {
    var data = JSON.parse(response);
    data.forEach(function(v) {
        var pattern = document.createElement("label");
        pattern.className = "pattern";
        pattern.innerText = v;
        pattern.onclick = async function() {
            await XHRrequest("DELETE", "/api/patterns/", JSON.stringify({pattern: v}));
            pattern.remove();
        }
        patternsList.appendChild(pattern);
    })
}

// -------------------------------- Variables -------------------------------- //

var patternsList = document.getElementById("patterns"),
    patternField = document.getElementById("pattern"); 

// ------------------------------ Functions ------------------------------ //

async function addPattern() {
    await XHRrequest("POST", "/api/patterns/", JSON.stringify({pattern: patternField.value}));
    location.reload();
}