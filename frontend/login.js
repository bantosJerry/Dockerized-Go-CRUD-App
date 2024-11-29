if (localStorage.getItem('authToken')) {
    window.location.href = 'home.html';
}

document.getElementById('loginForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;

    fetch('/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password })
    }).then(response => response.json())
      .then(data => {
          if (data.token) {
              localStorage.setItem('authToken', data.token);
              localStorage.setItem('userRole', data.user.role);
              localStorage.setItem('username', data.user.username);
              alert('Login successful');
              window.location.href = 'home.html';
          } else {
              alert('Invalid credentials');
          }
      }).catch(error => alert('Error logging in'));
});
