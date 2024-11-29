let currentUserId = null;

const userRole = localStorage.getItem('userRole') || 'User';
const username = localStorage.getItem('username') || 'Guest';
document.getElementById('userRole').textContent = userRole;
document.getElementById('username').textContent = username;

// Redirect to login if not authenticated
if (!localStorage.getItem('authToken')) {
    alert('You must be logged in to access this page');
    window.location.href = 'login.html';
}

// Logout function
document.getElementById('logoutButton').addEventListener('click', function () {
    localStorage.removeItem('authToken');
    localStorage.removeItem('userRole');
    localStorage.removeItem('username');
    window.location.href = 'login.html';
});

// Load all users and display in a table
document.getElementById('loadUsers').addEventListener('click', function () {
    const token = localStorage.getItem('authToken');
    if (!token) {
        alert('No authorization token found. Please log in.');
        return;
    }

    fetch('/api/users', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
        },
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(users => {
            const userTable = document.getElementById('userTable');
            const userTableBody = document.getElementById('userTableBody');
            userTableBody.innerHTML = '';
            userTable.style.display = 'table';

            users.forEach(user => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${user.id}</td>
                    <td>${user.username}</td>
                    <td>${user.email}</td>
                    <td>${user.role}</td>
                    <td>
                        <button class="update-button" onclick="showUpdateModal(${user.id}, '${user.username}', '${user.email}', '${user.role}')">Update</button>
                    </td>
                    <td>
                        <button class="delete-button" onclick="showDeleteModal(${user.id})">Delete</button>
                    </td>
                `;
                userTableBody.appendChild(row);
            });
        })
        .catch(error => {
            console.error('Error loading users:', error);
            alert('Error loading users. Please try again.');
        });
});

// Show update modal
function showUpdateModal(userId, username, email, role) {
    currentUserId = userId;
    document.getElementById('updateUsername').value = username;
    document.getElementById('updateEmail').value = email;
    document.getElementById('updateRole').value = role;
    document.getElementById('updateModal').style.display = 'block';
}

// Update user function
document.getElementById('updateUserButton').addEventListener('click', function () {
    const updatedUsername = document.getElementById('updateUsername').value;
    const updatedEmail = document.getElementById('updateEmail').value;
    const updatedRole = document.getElementById('updateRole').value;

    fetch(`/api/user/${currentUserId}`, {
        method: 'PUT',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('authToken')}`,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: updatedUsername,
            email: updatedEmail,
            role: updatedRole,
        }),
    })
        .then(response => response.json())
        .then(data => {
            alert('User updated successfully!');
            document.getElementById('updateModal').style.display = 'none';
            document.getElementById('loadUsers').click();
        })
        .catch(error => {
            alert('Error updating user.');
            console.error(error);
        });
});

// Show delete modal
function showDeleteModal(userId) {
    currentUserId = userId;
    document.getElementById('deleteModal').style.display = 'block';
}

// Delete user function
document.getElementById('deleteUserButton').addEventListener('click', function () {
    fetch(`/api/user/${currentUserId}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('authToken')}`,
            'Content-Type': 'application/json',
        },
    })
        .then(response => response.json())
        .then(data => {
            alert('User deleted successfully!');
            document.getElementById('deleteModal').style.display = 'none';
            document.getElementById('loadUsers').click();
        })
        .catch(error => {
            alert('Error deleting user.');
            console.error(error);
        });
});

// Close modals
document.getElementById('closeUpdateModal').addEventListener('click', function () {
    document.getElementById('updateModal').style.display = 'none';
});

document.getElementById('closeDeleteModal').addEventListener('click', function () {
    document.getElementById('deleteModal').style.display = 'none';
});
