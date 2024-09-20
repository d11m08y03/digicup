// darkmode.js

document.addEventListener('DOMContentLoaded', function() {
  const toggleButton = document.getElementById('dark-mode-toggle');
  const body = document.body;

  toggleButton.addEventListener('click', function() {
    // Toggle the 'dark' class on the <body> element
    body.classList.toggle('dark');

    // Save the preference in localStorage
    if (body.classList.contains('dark')) {
      localStorage.setItem('theme', 'dark');
    } else {
      localStorage.setItem('theme', 'light');
    }
  });

  // Load saved theme from localStorage
  if (localStorage.getItem('theme') === 'dark') {
    body.classList.add('dark');
  } else {
    body.classList.remove('dark');
  }
});
