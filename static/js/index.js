// Index page - Filter toggle and search
document.addEventListener('DOMContentLoaded', function() {
	// Toggle filters function
	window.toggleFilters = function() {
		const filters = document.getElementById('filtersSection') || document.getElementById('filters-section');
		const filterText = document.getElementById('filterText') || document.getElementById('filter-text');

		if (filters && filterText) {
			if (filters.classList.contains('hidden')) {
				filters.classList.remove('hidden');
				filterText.textContent = 'Скрыть фильтры';
			} else {
				filters.classList.add('hidden');
				filterText.textContent = 'Показать фильтры';
			}
		}
	};

	// Search judoka function
	window.searchJudoka = function() {
		const searchInput = document.getElementById('searchInput');
		if (searchInput) {
			const query = searchInput.value;
			console.log('Поиск:', query);
			// TODO: Implement search logic
		}
	};
});
