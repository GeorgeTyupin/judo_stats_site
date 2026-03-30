// Index page JS
document.addEventListener('DOMContentLoaded', function () {
	// Re-init Lucide icons after HTMX swaps
	document.body.addEventListener('htmx:afterSwap', function () {
		if (typeof lucide !== 'undefined') {
			lucide.createIcons();
		}
	});
});
