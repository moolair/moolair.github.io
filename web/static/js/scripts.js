document.addEventListener("DOMContentLoaded", () => {
    const postList = document.getElementById("post-list");

    fetch("/posts")
        .then(response => response.json())
        .then(posts => {
            if (!posts || posts.length === 0) {
                postList.innerHTML = "<p>No posts yet!</p>";
                return;
            }

            // 날짜 기준으로 정렬 (최신순)
            posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));

            posts.forEach(post => {
                const postElement = document.createElement("div");
                postElement.className = "post";
                postElement.innerHTML = `
                    <h3>${post.title || 'Untitled'}</h3>
                    <p>${post.content || 'No content'}</p>
                    <small>Posted: ${new Date(post.created_at).toLocaleString()}</small>
                `;
                postList.appendChild(postElement);
            });
        })
        .catch(error => {
            console.error('Error:', error);
            postList.innerHTML = `<p>Error loading posts: ${error}</p>`;
        });
});
