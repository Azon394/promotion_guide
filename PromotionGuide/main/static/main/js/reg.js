document.addEventListener('DOMContentLoaded', function() {
    var loginButton = document.querySelector('.reg-content_btn-auth');
    if (loginButton) {
        loginButton.addEventListener('click', function(event) {
            event.preventDefault();
            window.location.href = '/auth'; 
        });
    }
});
