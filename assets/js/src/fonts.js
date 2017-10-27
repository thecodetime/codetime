if (!sessionStorage.fontsLoaded) {
  var fontRennerMedium = new FontFaceObserver('renner_medium');
  var fontKanitLight = new FontFaceObserver('kanit_light');
  var fontKanitSemiBold = new FontFaceObserver('kanit_semibold');
  Promise.all([
    fontRennerMedium.load(),
    fontKanitLight.load(),
    fontKanitSemiBold.load()
  ]).then(function(){
    $('html').addClass('fonts-loaded');
    sessionStorage.fontsLoaded = true;
  });
} else {
  $('html').addClass('fonts-loaded');
}
