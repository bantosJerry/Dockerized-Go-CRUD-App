if (localStorage.getItem('authToken')) {
    window.location.href = 'home.html';
}

document.getElementById('registerForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const username = document.getElementById('registerUsername').value;
    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;
    const role = document.getElementById('role').value;

    fetch('/api/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, email, password, role })
    }).then(response => response.json())
      .then(data => {
          if (data.token) {
              alert('User registered successfully');
              window.location.href = 'login.html';
          } else {
              alert('Registration failed: ' + data.message);
          }
      }).catch(error => alert('Error registering user'));
});
