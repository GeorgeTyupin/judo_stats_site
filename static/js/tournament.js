// Tournament page - Gender filter for weight categories
document.addEventListener('DOMContentLoaded', function() {
	const filterButtons = document.querySelectorAll('.filter-btn');
	const genderSections = document.getElementById('gender-sections');
	const menSection = document.querySelector('[data-gender="men"]');
	const womenSection = document.querySelector('[data-gender="women"]');

	filterButtons.forEach(button => {
		button.addEventListener('click', function() {
			const filter = this.getAttribute('data-filter');

			// Update active state
			filterButtons.forEach(btn => btn.classList.remove('active'));
			this.classList.add('active');

			// Apply filter with animation
			if (filter === 'all') {
				genderSections.classList.remove('single-column');
				menSection.classList.remove('hidden');
				womenSection.classList.remove('hidden');
			} else if (filter === 'men') {
				genderSections.classList.add('single-column');
				menSection.classList.remove('hidden');
				womenSection.classList.add('hidden');
			} else if (filter === 'women') {
				genderSections.classList.add('single-column');
				menSection.classList.add('hidden');
				womenSection.classList.remove('hidden');
			}
		});
	});
});
