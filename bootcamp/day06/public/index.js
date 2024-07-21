console.log("Script loaded")

const FULL_SCREEN_STYLES = {
	position: "fixed",
	top: 0,
	left: 0,
	width: "100vw",
	height: "100vh",
	zIndex: '1000',
	backgroundColor: 'black',
	maxWidth: "100%"
}

const FULL_SCREEN_ATTRIBUTE_NAME = "data-fullscreen";

/**
* @param { HTMLElement } alert
*/
function removeAlert(alert) {
	/** @type {HTMLElement} parent */
	alert.classList.add("alert-remove")
	setTimeout(() => {
		alert.remove()
	}, 300)
}


/**
* @param {MouseEvent} ev 
*/
function cancelAlert(ev) {
	/** @type {HTMLElement} parent */
	const parent = ev.target.offsetParent;
	removeAlert(parent)
}

function newAlertCleaningWorker() {
	const perAlertDuration = 10000;
	const mainLoopDuration = 10500;
	let cleaningLoopTimeoutId;
	function stopLoop() {
		console.log("LOOP STOP")
		clearTimeout(cleaningLoopTimeoutId)
	}
	function continueLoop() {
		console.log("LOOP CONTINUE")
		alertCleaningLoop()
	}
	const alertsContainer = document.getElementById("alerts")
	if (!alertsContainer) {
		return
	}
	alertsContainer.addEventListener("mouseenter", stopLoop)
	alertsContainer.addEventListener("mouseleave", continueLoop)

	/**
	 * @param {HTMLElement} toBeRemoved 
	 */
	function scheduleRemoveAlert(toBeRemoved) {
		let countDownProgress = toBeRemoved.getElementsByClassName("alert-countdown")[0]
		countDownProgress.classList.add("active")
		toBeRemovedTimeoutId = setTimeout(() => {
			removeAlert(toBeRemoved)
		}, perAlertDuration)
	}

	console.log("initialized")
	function alertCleaningLoop() {
		cleaningLoopTimeoutId = setTimeout(alertCleaningLoop, mainLoopDuration)
		console.log("LOOP RUN")
		if (alertsContainer.children.length == 0) {
			return
		}
		scheduleRemoveAlert(alertsContainer.children[0])
	}
	alertCleaningLoop()
}

newAlertCleaningWorker()

console.log(htmx.config.useTemplateFragments)
htmx.config.useTemplateFragments = true;
console.log(htmx.config.useTemplateFragments)

/** Toggles article form full screen
 * @param {string} id - id of the element to toggle full screen
 * */
function toggleFullScreenById(id) {
	const form = document.getElementById(id)
	if (form) {
		if (form.getAttribute(FULL_SCREEN_ATTRIBUTE_NAME)) {
			form.style = null;
			form.removeAttribute(FULL_SCREEN_ATTRIBUTE_NAME);
			return
		}
		Object.entries(FULL_SCREEN_STYLES).map(([key, val]) => {
			form.style[key] = val;
		})
		form.setAttribute(FULL_SCREEN_ATTRIBUTE_NAME, "true")
	}
}

