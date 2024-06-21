// Get the modal
var modal = document.getElementById("idLogin");
var modal2 = document.getElementById("idSignup");

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == modal) {
        modal.style.display = "none";
    } else if (event.target == modal2) {
        modal2.style.display = "none";
    }
};

const close_modal = (id) => {
    document.getElementById(id).style.display = "none";
    clear_login();
    clear_signup();
};

const clear_login = () => {
    document.getElementById("usernameidLogin").value = "";
    document.getElementById("passwordidLogin").value = "";
};

const clear_signup = () => {
    document.getElementById("usernameidSignup").value = "";
    document.getElementById("passwordidSignup").value = "";
    document.getElementById("emailidSignup").value = "";
};

const show_password = (id) => {
    let password = document.getElementById("password"+id);
    if (password.type === "password") {
        password.type = "text";
    } else {
        password.type = "password";
    }
};