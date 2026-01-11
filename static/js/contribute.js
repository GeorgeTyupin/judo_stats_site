// Contribute page - Form type selection and reset
function selectType(type) {
	// Update type cards
	document.querySelectorAll('.type-card').forEach(card => {
		card.classList.remove('active');
	});
	event.currentTarget.classList.add('active');

	// Hide all forms
	document.querySelectorAll('.form-section').forEach(form => {
		form.classList.remove('active');
	});

	// Show selected form
	document.getElementById(type + '-form').classList.add('active');

	// Scroll to form
	setTimeout(() => {
		document.getElementById(type + '-form').scrollIntoView({
			behavior: 'smooth',
			block: 'start'
		});
	}, 100);
}

function resetForm(type) {
	const form = document.querySelector('#' + type + '-form form');
	if (form && confirm('Вы уверены, что хотите очистить форму?')) {
		form.reset();
	}
}
