// Tournament page — gender filter, champions filter, tabs
document.addEventListener('DOMContentLoaded', function() {

	// ── Gender / Champions filter ─────────────────────────────────────────────
	const filterButtons = document.querySelectorAll('.filter-btn');
	const genderSections = document.getElementById('gender-sections');
	const menSection = document.querySelector('[data-gender="men"]');
	const womenSection = document.querySelector('[data-gender="women"]');

	function showAllResults() {
		document.querySelectorAll('.result-item').forEach(item => {
			item.style.display = '';
		});
	}

	function showChampionsOnly() {
		document.querySelectorAll('.result-item').forEach(item => {
			item.style.display = item.getAttribute('data-place') === '1' ? '' : 'none';
		});
	}

	filterButtons.forEach(button => {
		button.addEventListener('click', function() {
			const filter = this.getAttribute('data-filter');

			filterButtons.forEach(btn => btn.classList.remove('active'));
			this.classList.add('active');

			if (filter === 'champions') {
				genderSections.classList.remove('single-column');
				menSection.classList.remove('hidden');
				womenSection.classList.remove('hidden');
				showChampionsOnly();
			} else {
				showAllResults();
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
			}
		});
	});

	// ── Tabs ──────────────────────────────────────────────────────────────────
	const tabButtons = document.querySelectorAll('.tab-btn');
	const tabContents = document.querySelectorAll('.tab-content');

	tabButtons.forEach(button => {
		button.addEventListener('click', function() {
			const targetTab = this.getAttribute('data-tab');

			tabButtons.forEach(btn => btn.classList.remove('active'));
			this.classList.add('active');

			tabContents.forEach(content => {
				if (content.getAttribute('data-tab-content') === targetTab) {
					content.classList.remove('tab-content-hidden');
				} else {
					content.classList.add('tab-content-hidden');
				}
			});
		});
	});
});
