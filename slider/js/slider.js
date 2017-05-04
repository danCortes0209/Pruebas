$(function() {
    var SliderModule = (function() {
        var pb = {};
        pb.el = $('#slider');
        pb.items = {
            panel: pb.el.find('li')
        }
        
        var SliderInterval,
            currentSlider = 0,
            nextSlider = 1,
            lengthSlider = pb.items.panel.length;
        
        pb.init = function(settings){
            var output = '';
            
            SliderInit();
            
            for(var i = 0;i<lengthSlider;i++){
                if(i==0){
                    output += '<li class="active"></li>';
                } else {
                    output += '<li></li>';
                }
            }
            $('#controls').html(output).on('click','li',function(e){
                var $this = $(this);
                if(currentSlider !== $this.index()){
                changePanel($this.index());                 
                }
            });
        }
        
        var SliderInit = function(){
            SliderInterval = setInterval(pb.startSlider, 3000);
        }
        
        pb.startSlider = function(){
            var panels = pb.items.panel,
                controls = $('#controls li');
            if(nextSlider >= lengthSlider){
                nextSlider = 0;
                currentSlider = lengthSlider-1;
            }
            controls.removeClass('active').eq(currentSlider).addClass('active');
            panels.eq(currentSlider).fadeOut('slow');
            panels.eq(nextSlider).fadeIn('slow');
            
            currentSlider = nextSlider;
            nextSlider += 1;
            
        }
        
        var changePanel = function(id){
            clearInterval(SliderInterval);
            var panels = pb.items.panel, controls = $('#controls li');
            if (id >= lengthSlider){
                id = 0;
            } else if (id < 0){
                id = lengthSlider-1;
            }
            controls.removeClass('active').eq(id).addClass('active');
            panels.eq(currentSlider).fadeOut('slow');
            panels.eq(id).fadeIn('slow');
            
            currentSlider = id;
            nextSlider = id+1;
            
            SliderInit();
        }
        
        return pb;
    }());
    SliderModule.init();
});