function AJAXrequest(method, domain, data) {
    $.ajax({
        type: method,
        url: domain,
        data: data,
        success: function() {
            window.location.reload();
        },
    });
}
async function addProduct() {
    await AJAXrequest("POST", "/api/product/add/", JSON.stringify({
        name: document.getElementById("name").value,
        shop_url: document.getElementById("shop_url").value,
        link: document.getElementById("link").value,
        photo_url: document.getElementById("photo_url").value,
        type: "new arrival",
    }));
}