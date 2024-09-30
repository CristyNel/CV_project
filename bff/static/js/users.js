function updateUser(userId) {
    fetch(`/user/${userId}`)
    .then(res => {console.log(res)})
    .catch(error => {console.log(error)});
}

function deleteUser(userId) {
    fetch(`/user?id=${userId}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(errorData => {
                throw new Error(errorData.error || 'Network response was not ok');
            });
        }
        return response.json();
    })
    .then((data) => {
        if (data.success) {
            alert('User deleted successfully');
            window.location.reload();
        } else {
            throw new Error(data.error || 'User deletion failed');
        }
    })
    .catch((error) => {
        alert(`Error deleting user: ${error.message}`);
        console.error('There was a problem with the fetch operation:', error.message);
    });
}