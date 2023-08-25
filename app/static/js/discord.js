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
        },
        error: function() {
            console.log("An error occured while loading...");
        }
    });
}

// -------------------------------- Setup -------------------------------- //

window.onload = async function() {
    var response = await XHRrequest("GET", "/api/discord/");
    console.log(JSON.parse(response));
}

