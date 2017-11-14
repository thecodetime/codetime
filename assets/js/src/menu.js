$(function(){
  $mobileMenu = $('#mobile-menu');
  $('#mobile-menu-toggle').click(function(){
    if ($mobileMenu.is(':visible')) {
      $(this).removeClass('selected');
      $mobileMenu.slideUp('slow');
    } else {
      $(this).addClass('selected');
      $mobileMenu.slideDown('slow');
    }
  });
});
