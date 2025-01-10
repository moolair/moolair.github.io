document.addEventListener("DOMContentLoaded", () => {
    const postList = document.getElementById("post-list");

    // Fetch posts from the API
    fetch("/posts")
        .then(response => response.json())
        .then(posts => {
            posts.forEach(post => {
                const postElement = document.createElement("div");
                postElement.className = "post";
                postElement.innerHTML = `
                    <h3>${post.title}</h3>
                    <p>${post.content}</p>
                `;
                postList.appendChild(postElement);
            });
        })
        .catch(error => {
            postList.innerHTML = `<p>Error loading posts: ${error}</p>`;
        });
});
