// Index page - Filter toggle and search
document.addEventListener('DOMContentLoaded', function() {
	// Утилита для плавного slide toggle с Web Animations API
	function slideToggle(element, duration = 500) {
		const isHidden = element.classList.contains('hidden');

		if (isHidden) {
			// Разворачивание
			element.classList.remove('hidden');
			element.style.overflow = 'hidden';

			requestAnimationFrame(() => {
				// Используем getBoundingClientRect для точного измерения с учетом margin
				const endHeight = element.getBoundingClientRect().height;

				// Устанавливаем начальную высоту
				element.style.height = '0px';

				requestAnimationFrame(() => {
					const animation = element.animate([
						{ height: '0px' },
						{ height: endHeight + 'px' }
					], {
						duration: duration,
						easing: 'cubic-bezier(0.4, 0, 0.2, 1)'
					});

					animation.onfinish = () => {
						element.style.height = '';
						element.style.overflow = '';
					};
				});
			});
		} else {
			// Сворачивание
			// Измеряем текущую высоту с помощью getBoundingClientRect
			const startHeight = element.getBoundingClientRect().height;

			element.style.overflow = 'hidden';
			element.style.height = startHeight + 'px';

			requestAnimationFrame(() => {
				requestAnimationFrame(() => {
					const animation = element.animate([
						{ height: startHeight + 'px' },
						{ height: '0px' }
					], {
						duration: duration,
						easing: 'cubic-bezier(0.4, 0, 0.2, 1)'
					});

					animation.onfinish = () => {
						element.classList.add('hidden');
						element.style.height = '';
						element.style.overflow = '';
					};
				});
			});
		}
	}

	// Toggle filters function
	window.toggleFilters = function() {
		const filtersWrapper = document.getElementById('filters-wrapper');
		const filterText = document.getElementById('filterText');
		const toggleButton = document.querySelector('.toggle-filters');

		if (filtersWrapper && filterText && toggleButton) {
			const isHidden = filtersWrapper.classList.contains('hidden');

			// Анимируем с помощью Web Animations API
			slideToggle(filtersWrapper, 500);

			// Обновляем UI кнопки
			if (isHidden) {
				toggleButton.classList.remove('collapsed');
				filterText.textContent = 'Скрыть фильтры';
			} else {
				toggleButton.classList.add('collapsed');
				filterText.textContent = 'Показать фильтры';
			}
		}
	};

	// Set active category - вызывается после HTMX загрузки фильтров
	document.body.addEventListener('htmx:afterSwap', function(event) {
		if (event.detail.target.id === 'search-filters') {
			// Триггерим поиск после смены категории, если есть текст
			const searchInput = document.getElementById('searchInput');
			if (searchInput && searchInput.value) {
				htmx.trigger(searchInput, 'keyup');
			}
		}
	});
});

// Set active category - вызывается при клике
function setActiveCategory(button) {
	// Удаляем активный класс у всех кнопок
	const allButtons = document.querySelectorAll('.category-tab');
	allButtons.forEach(btn => {
		btn.classList.remove('active');
		btn.style.background = '#f0f4f8';
		btn.style.color = '#0F172A';
	});

	// Добавляем активный класс к выбранной кнопке
	button.classList.add('active');
	button.style.background = 'linear-gradient(135deg, #14B8A6 0%, #0D9488 100%)';
	button.style.color = 'white';
}
