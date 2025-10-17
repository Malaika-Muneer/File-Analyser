// ===================== SIGN UP / SIGN IN / UPLOAD =====================
document.addEventListener("DOMContentLoaded", () => {
  const signupForm = document.getElementById("signupForm");
  const signinForm = document.getElementById("signinForm");
  const uploadForm = document.getElementById("uploadForm");

  // ===================== SIGN UP =====================
  if (signupForm) {
    signupForm.addEventListener("submit", async (e) => {
      e.preventDefault();

      const username = document.getElementById("username").value.trim();
      const email = document.getElementById("email").value.trim();
      const password = document.getElementById("password").value.trim();

      try {
        const res = await fetch("/signup", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ username, email, password }),
        });

        const data = await res.json();
        const msg = document.getElementById("signupMsg");

        if (res.ok) {
          msg.style.color = "green";
          msg.textContent = data.message || "Signup successful! Redirecting...";
          setTimeout(() => {
            window.location.href = "/signin.html";
          }, 1500);
        } else {
          msg.style.color = "red";
          msg.textContent = data.error || "Signup failed!";
        }
      } catch (error) {
        console.error("Signup error:", error);
      }
    });
  }

  // ===================== SIGN IN =====================
  if (signinForm) {
    signinForm.addEventListener("submit", async (e) => {
      e.preventDefault();

      const username = document.getElementById("username").value.trim();
      const password = document.getElementById("password").value.trim();

      try {
        const res = await fetch("/signin", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ username, password }),
        });

        const data = await res.json();
        const msg = document.getElementById("signinMsg");

        if (res.ok && data.token) {
          localStorage.setItem("jwtToken", data.token);
          msg.style.color = "green";
          msg.textContent = "Login successful! Redirecting...";
          setTimeout(() => {
            window.location.href = "/upload.html";
          }, 1500);
        } else {
          msg.style.color = "red";
          msg.textContent = data.error || "Invalid credentials!";
        }
      } catch (error) {
        console.error("Signin error:", error);
      }
    });
  }

  // ===================== UPLOAD =====================
  if (uploadForm) {
    uploadForm.addEventListener("submit", async (e) => {
      e.preventDefault();

      const token = localStorage.getItem("jwtToken");
      const file = document.getElementById("fileInput").files[0];
      if (!file) return alert("Please select a file");

      const formData = new FormData();
      formData.append("file", file);

      try {
        const res = await fetch("/upload", {
          method: "POST",
          headers: { Authorization: `Bearer ${token}` },
          body: formData,
        });

        const data = await res.json();
        const uploadResults = document.getElementById("uploadResults");
        uploadResults.innerHTML = "";

        if (res.ok) {
          uploadResults.textContent = JSON.stringify(data, null, 4);
        } else {
          uploadResults.textContent = JSON.stringify(data, null, 4) || "Error uploading file";
        }
      } catch (error) {
        console.error("Upload error:", error);
      }
    });
  }
}); // ✅ This was missing earlier
