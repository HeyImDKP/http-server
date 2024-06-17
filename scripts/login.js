// Get the modal
var modal = document.getElementById("id01");

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
};

const close_modal = () => {
    document.getElementById("id01").style.display = "none";
    clear_login();
};

const clear_login = () => {
    document.getElementById("username").value = "";
    document.getElementById("password").value = "";
};

const show_password = () => {
    let password = document.getElementById("password");
    if (password.type === "password") {
        password.type = "text";
    } else {
        password.type = "password";
    }
};