<<<<<<< HEAD
$(function() {

  var
    toc = $('#toc').show(),
    items = $('#toc > ul').hide();

  $('#toc h3')
    .click(function() {
      if (items.is(':visible')) {
        items.animate({
          height:     'hide',
          opacity:    'hide'
        }, 300, function() {
          toc.removeClass('expandedtoc');
        });
      }
      else {
        items.animate({
          height:     'show',
          opacity:    'show'
        }, 400);
        toc.addClass('expandedtoc');
      }
    });

});
=======
$(function() {

  var
    toc = $('#toc').show(),
    items = $('#toc > ul').hide();

  $('#toc h3')
    .click(function() {
      if (items.is(':visible')) {
        items.animate({
          height:     'hide',
          opacity:    'hide'
        }, 300, function() {
          toc.removeClass('expandedtoc');
        });
      }
      else {
        items.animate({
          height:     'show',
          opacity:    'show'
        }, 400);
        toc.addClass('expandedtoc');
      }
    });

});
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
