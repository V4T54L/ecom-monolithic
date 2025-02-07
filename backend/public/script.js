const signupForm = document.getElementById('signupForm');
const loginForm = document.getElementById('loginForm');
const userProfile = document.getElementById('userProfile');
const formTitle = document.getElementById('formTitle');
const profileUsername = document.getElementById('profileUsername');
const toggleToLogin = document.getElementById('toggleToLogin');
const toggleToSignup = document.getElementById('toggleToSignup');
const logoutButton = document.getElementById('logoutButton');
const authPage = document.getElementById('auth-page');

const TOKEN_KEY = 'authToken';

// Prioritize fetch functions for handling relevant API requests
async function createUser(username, password) {
    try {
        const response = await axios.post('/auth/signup', { username, password });
        alert('User created successfully!');
        toggleForms();
    } catch (error) {
        alert(error.response?.data.error || error.message);
    }
}

async function login(username, password) {
    try {
        const response = await axios.post('/auth/login', { username, password });
        localStorage.setItem(TOKEN_KEY, response.data.token);
        fetchUserProfile();
    } catch (error) {
        alert(error.response?.data.error || error.message);
    }
}

async function fetchUserProfile() {
    const token = localStorage.getItem(TOKEN_KEY);
    if (!token) {
        return;
    }

    try {
        const response = await axios.get('/user/profile', {
            headers: { Authorization: `Bearer ${token}` }
        });
        const user = response.data.user;
        if (user) {
            profileUsername.textContent = `Welcome, ${user.username}!`;
            userProfile.classList.remove('hidden');
            authPage.classList.add('hidden');
        }
    } catch (error) {
        alert(error.response?.data.error || error.message);
        localStorage.removeItem(TOKEN_KEY);
        toggleForms(false);
        authPage.classList.remove('hidden');
    }
}

function toggleForms(isLoginForm = true) {
    if (isLoginForm) {
        formTitle.textContent = "Sign Up";
        signupForm.classList.toggle('hidden');
        loginForm.classList.toggle('hidden');
        userProfile.classList.add('hidden');
    } else {
        formTitle.textContent = "Login";
        loginForm.classList.toggle('hidden');
        signupForm.classList.toggle('hidden');
        userProfile.classList.add('hidden');
    }
}

// Event Listeners
signupForm.addEventListener('submit', function (e) {
    e.preventDefault();
    const username = document.getElementById('signupUsername').value;
    const password = document.getElementById('signupPassword').value;
    createUser(username, password);
});

loginForm.addEventListener('submit', function (e) {
    e.preventDefault();
    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;
    login(username, password);
});

toggleToLogin.addEventListener('click', () => toggleForms(false));
toggleToSignup.addEventListener('click', () => toggleForms(true));

logoutButton.addEventListener('click', () => {
    localStorage.removeItem(TOKEN_KEY);
    userProfile.classList.add('hidden');
    authPage.classList.remove('hidden');
    toggleForms(true);
});

// Check user profile on page load
document.addEventListener('DOMContentLoaded', fetchUserProfile);