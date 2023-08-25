function AJAXrequest(method, domain, data, redirectURL) {
    $.ajax({
        type: method,
        url: domain,
        data: data,
        success: function() {
            window.location = redirectURL;
        },
    });
}
async function login() {
    await AJAXrequest("POST", "/api/auth/", JSON.stringify({
        username: document.getElementById("username").value,
        password: document.getElementById("password").value,
    }), "/products");
}

async function join() {
    await AJAXrequest("POST", "/api/join/", JSON.stringify({
        username: document.getElementById("username").value,
        password: document.getElementById("password").value,
        invite_key: document.getElementById("invite_key").value,
    }), "/auth");
}