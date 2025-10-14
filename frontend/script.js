const API_BASE_URL = ""; // same origin


// --- SIGNUP ---
const signupForm = document.getElementById("signupForm");
if (signupForm) {
  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const user = {
      username: document.getElementById("username").value,
      email: document.getElementById("email").value,
      password: document.getElementById("password").value,
    };

    try {
      const res = await fetch(`${API_BASE_URL}/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(user),
      });

      if (!res.ok) {
        const errData = await res.json();
        document.getElementById("signupMsg").innerText =
          errData.error || "Signup failed!";
        return;
      }

      const data = await res.json();
      document.getElementById("signupMsg").innerText =
        data.message || "Signup successful!";
    } catch (err) {
      console.error("Signup error:", err);
      document.getElementById("signupMsg").innerText =
        "Error connecting to server!";
    }
  });
}

// --- SIGNIN ---
const signinForm = document.getElementById("signinForm");
if (signinForm) {
  signinForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const creds = {
      username: document.getElementById("username").value,
      password: document.getElementById("password").value,
    };

    try {
      const res = await fetch(`${API_BASE_URL}/signin`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(creds),
      });

      const data = await res.json();
      if (res.ok && data.token) {
        localStorage.setItem("jwtToken", data.token);
        document.getElementById("signinMsg").innerText = "Login successful!";
        // Redirect to upload page after a short delay
        setTimeout(() => (window.location.href = "upload.html"), 1000);
      } else {
        document.getElementById("signinMsg").innerText =
          data.error || "Login failed!";
      }
    } catch (err) {
      console.error("Signin error:", err);
      document.getElementById("signinMsg").innerText =
        "Error connecting to server!";
    }
  });
}

// --- UPLOAD ---
const uploadForm = document.getElementById("uploadForm");
if (uploadForm) {
  uploadForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const token = localStorage.getItem("jwtToken");
    if (!token) {
      document.getElementById("uploadMsg").innerText = "Please sign in first!";
      return;
    }

    const fileInput = document.getElementById("fileInput");
    if (!fileInput.files[0]) {
      document.getElementById("uploadMsg").innerText = "Please select a file!";
      return;
    }

    const formData = new FormData();
    formData.append("file", fileInput.files[0]);

    try {
      const res = await fetch(`${API_BASE_URL}/upload`, {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
        body: formData,
      });

      const data = await res.json();
      if (res.ok) {
        document.getElementById("uploadMsg").innerText =
          JSON.stringify(data, null, 2); // show full file analysis
      } else {
        document.getElementById("uploadMsg").innerText =
          data.error || "File upload failed!";
      }
    } catch (err) {
      console.error("Upload error:", err);
      document.getElementById("uploadMsg").innerText =
        "Error connecting to server!";
    }
  });
}
