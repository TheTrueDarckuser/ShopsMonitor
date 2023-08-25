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

// -------------------------------- Setup -------------------------------- //

window.onload = async function() {
    var response = await XHRrequest("GET", "/api/products/", null);
    var data = JSON.parse(response);
    applyData(data);
}

// -------------------------------- Variables -------------------------------- //

var productsList = document.getElementById("products-list");

// ------------------------------ Functions ------------------------------ //

function applyData(data) {
    var n = (!data?0:data.length);
    updateDivs(n);

    var i = 0;
    data.forEach(function(v) {
        var pWrapper = productsList.childNodes[i++];
        pWrapper.childNodes[0].href = v["link"]
        pWrapper.childNodes[0].childNodes[0].src = v["photo_url"]
        pWrapper.childNodes[0].childNodes[0].alt = v["name"]
        pWrapper.childNodes[1].innerText = v["name"]
        pWrapper.childNodes[2].childNodes[0].onclick = async function() {
            await XHRrequest("DELETE", "/api/products/", JSON.stringify({link: v["link"]}));
            pWrapper.remove();
        }
    });
}

function updateDivs(n) {
    for(var i=0; i<n; i++) {
        product = document.createElement("li");
        product.className = "product";

        productLink = document.createElement("a");
        productLink.className = "product-link";
        productLink.target = "_blank";

        productButtons = document.createElement("div");
        productButtons.className = "control-buttons";

        productButtonIgnore = document.createElement("button");
        productButtonIgnore.className = "button-ignore";
        productButtonIgnore.innerText = "Ignore";
        productButtonIgnore.type = "button";

        productImage = document.createElement("img");
        productImage.className = "product-image";

        productName = document.createElement("div");
        productName.className = "product-name";

        productsList.appendChild(product);
        product.appendChild(productLink);
        productLink.appendChild(productImage);
        product.appendChild(productName);
        product.appendChild(productButtons);
        productButtons.appendChild(productButtonIgnore);
    }
}

// ---------------------- Search Bar ---------------------- //

async function search() {
    var input = document.getElementById("search-bar").value;

    await setProductsDefault();
    productsList.childNodes.forEach(function(v){
        if(!v.childNodes[1].innerText.toLowerCase().includes(input.toLowerCase())) {
            v.style.display = "none";
        }
    });
}

async function setProductsDefault() {
    productsList.childNodes.forEach(function(v) {
        v.style.display = "unset";
    });
}