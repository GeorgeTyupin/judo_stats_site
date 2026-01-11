// Judoka page - Tab switching
function openTab(tabName) {
	// Hide all tab contents
	const tabContents = document.getElementsByClassName('tab-content');
	for (let i = 0; i < tabContents.length; i++) {
		tabContents[i].classList.remove('active');
	}

	// Remove active class from all tabs
	const tabs = document.getElementsByClassName('tab');
	for (let i = 0; i < tabs.length; i++) {
		tabs[i].classList.remove('active');
	}

	// Show the selected tab
	document.getElementById(tabName).classList.add('active');
	event.target.classList.add('active');
}
