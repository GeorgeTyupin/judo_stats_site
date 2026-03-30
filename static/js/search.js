// Search page JS — sort, row navigation, filter chips, autocomplete skeleton, toggleFilters

document.addEventListener('DOMContentLoaded', function () {

	// ============================================================
	// Slide toggle (Web Animations API)
	// ============================================================
	function slideToggle(element, duration) {
		duration = duration || 500;
		var isHidden = element.classList.contains('hidden');

		if (isHidden) {
			element.classList.remove('hidden');
			element.style.overflow = 'hidden';

			requestAnimationFrame(function () {
				var endHeight = element.getBoundingClientRect().height;
				element.style.height = '0px';

				requestAnimationFrame(function () {
					var anim = element.animate(
						[{ height: '0px' }, { height: endHeight + 'px' }],
						{ duration: duration, easing: 'cubic-bezier(0.4, 0, 0.2, 1)' }
					);
					anim.onfinish = function () {
						element.style.height = '';
						element.style.overflow = '';
					};
				});
			});
		} else {
			var startHeight = element.getBoundingClientRect().height;
			element.style.overflow = 'hidden';
			element.style.height = startHeight + 'px';

			requestAnimationFrame(function () {
				requestAnimationFrame(function () {
					var anim = element.animate(
						[{ height: startHeight + 'px' }, { height: '0px' }],
						{ duration: duration, easing: 'cubic-bezier(0.4, 0, 0.2, 1)' }
					);
					anim.onfinish = function () {
						element.classList.add('hidden');
						element.style.height = '';
						element.style.overflow = '';
					};
				});
			});
		}
	}

	// ============================================================
	// Toggle filters panel
	// ============================================================
	window.toggleFilters = function () {
		var wrapper = document.getElementById('filters-wrapper');
		var filterText = document.getElementById('filterText');
		var btn = document.querySelector('.toggle-filters');
		if (!wrapper) return;

		var isHidden = wrapper.classList.contains('hidden');
		slideToggle(wrapper, 500);

		if (btn) {
			if (isHidden) {
				btn.classList.remove('collapsed');
			} else {
				btn.classList.add('collapsed');
			}
		}
		if (filterText) {
			filterText.textContent = isHidden ? 'Скрыть фильтры' : 'Показать фильтры';
		}
	};

	// ============================================================
	// Active category tab
	// ============================================================
	window.setActiveCategory = function (button) {
		document.querySelectorAll('.category-tab').forEach(function (btn) {
			btn.classList.remove('active');
		});
		button.classList.add('active');
	};

	// ============================================================
	// Sorting
	// ============================================================
	function triggerSort(col, dir) {
		var sortByInput = document.getElementById('sortBy');
		var sortDirInput = document.getElementById('sortDir');
		if (sortByInput) sortByInput.value = col;
		if (sortDirInput) sortDirInput.value = dir;

		// Trigger search via the search button (which has HTMX attrs)
		var searchBtn = document.getElementById('searchBtn');
		if (searchBtn) {
			htmx.trigger(searchBtn, 'click');
		}
	}

	function attachTableSortHandlers(container) {
		container = container || document;
		container.querySelectorAll('.sort-bar-btn').forEach(function (btn) {
			btn.addEventListener('click', function () {
				var col = this.dataset.sortCol;
				var dir = this.dataset.sortDir || 'asc';
				triggerSort(col, dir);
			});
		});
	}

	function updateSortBarActive() {
		var sortByInput = document.getElementById('sortBy');
		var currentCol = sortByInput ? sortByInput.value : '';
		document.querySelectorAll('.sort-bar-btn').forEach(function (btn) {
			if (btn.dataset.sortCol === currentCol) {
				btn.classList.add('sort-active-btn');
			} else {
				btn.classList.remove('sort-active-btn');
			}
		});
	}

	function attachRowNavHandlers(container) {
		container = container || document;
		container.querySelectorAll('.result-row[data-href]').forEach(function (row) {
			row.addEventListener('click', function () {
				window.location.href = this.dataset.href;
			});
		});
	}

	// ============================================================
	// Filter chips
	// ============================================================
	var filterLabels = {
		'filter_country':     'Страна',
		'filter_gender':      'Пол',
		'filter_birth_year':  'Год рождения',
		'filter_sport_club':  'Спортивное общество',
		'filter_city':        'Город',
		'filter_type':        'Тип',
		'filter_year_from':   'Год от',
		'filter_year_to':     'Год до',
		'filter_month_from':  'Месяц от',
		'filter_month_to':    'Месяц до',
		'filter_republic':    'Республика',
		'filter_age_group':   'Возраст. группа',
		'filter_oblast':      'Область',
	};

	function updateFilterChips() {
		var chipsContainer = document.getElementById('filter-chips');
		var activeFiltersBlock = document.getElementById('active-filters');
		if (!chipsContainer || !activeFiltersBlock) return;

		chipsContainer.innerHTML = '';
		var hasActive = false;

		document.querySelectorAll('[name^="filter_"]').forEach(function (input) {
			var val = input.value.trim();
			if (!val) return;
			hasActive = true;

			var label = filterLabels[input.name] || input.name;
			var chip = document.createElement('span');
			chip.className = 'filter-chip';
			chip.innerHTML =
				'<span>' + label + ': ' + escapeHtml(val) + '</span>' +
				'<button class="filter-chip-remove" type="button" data-target="' + input.name + '">' +
					'<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>' +
				'</button>';
			chipsContainer.appendChild(chip);
		});

		activeFiltersBlock.style.display = hasActive ? 'flex' : 'none';
	}

	function escapeHtml(str) {
		var div = document.createElement('div');
		div.appendChild(document.createTextNode(str));
		return div.innerHTML;
	}

	// Chip remove click — event delegation
	document.addEventListener('click', function (e) {
		var btn = e.target.closest('.filter-chip-remove');
		if (!btn) return;
		var targetName = btn.dataset.target;
		var input = document.querySelector('[name="' + targetName + '"]');
		if (input) {
			input.value = '';
			updateFilterChips();
			triggerSearch();
		}
	});

	// Clear all filters
	window.clearAllFilters = function () {
		document.querySelectorAll('[name^="filter_"]').forEach(function (input) {
			input.value = '';
		});
		updateFilterChips();
		triggerSearch();
	};

	function triggerSearch() {
		var searchBtn = document.getElementById('searchBtn');
		if (searchBtn) htmx.trigger(searchBtn, 'click');
	}

	// Update chips whenever any filter input changes
	document.addEventListener('change', function (e) {
		if (e.target && e.target.name && e.target.name.startsWith('filter_')) {
			updateFilterChips();
		}
	});
	document.addEventListener('input', function (e) {
		if (e.target && e.target.name && e.target.name.startsWith('filter_')) {
			updateFilterChips();
		}
	});

	// ============================================================
	// Autocomplete skeleton (UI only, no real API yet)
	// ============================================================
	document.querySelectorAll('.autocomplete-wrapper').forEach(function (wrapper) {
		var input = wrapper.querySelector('.autocomplete-input');
		var dropdown = wrapper.querySelector('.autocomplete-dropdown');
		if (!input || !dropdown) return;

		input.addEventListener('focus', function () {
			dropdown.removeAttribute('hidden');
		});

		input.addEventListener('blur', function () {
			// Delay so clicks inside dropdown register first
			setTimeout(function () {
				dropdown.setAttribute('hidden', '');
			}, 150);
		});
	});

	// ============================================================
	// HTMX afterSwap — reinit Lucide, table handlers, chips
	// ============================================================
	document.body.addEventListener('htmx:afterSwap', function (event) {
		if (typeof lucide !== 'undefined') {
			lucide.createIcons();
		}

		var target = event.detail.target;
		attachTableSortHandlers(target);
		attachRowNavHandlers(target);
		updateFilterChips();
		updateSortBarActive();

		// After category swap: trigger search if query present
		if (target.id === 'search-filters') {
			var searchInput = document.getElementById('searchInput');
			if (searchInput && searchInput.value) {
				triggerSearch();
			}
		}
	});

	// Initial attach (for SSR-rendered results)
	attachTableSortHandlers();
	attachRowNavHandlers();
	updateFilterChips();
	updateSortBarActive();
});
