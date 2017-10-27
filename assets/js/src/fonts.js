if (!sessionStorage.fontsLoaded) {
  var fontRennerMedium = new FontFaceObserver('renner_medium');
  fontRennerMedium.load().then(function() {
    $('html').addClass('fonts-loaded');
    sessionStorage.fontsLoaded = true;
  });
} else {
  $('html').addClass('fonts-loaded');
}
