$(function() {
    $('.js-change').change(function() {
	console.log("oh snap something changed");
	console.log(getValues());
    });
    
});

function getValues() {
    return {
	Orbit: $('#orbit').val(),
	Zoom: $('#zoom').val(),
	PresetValue: $("input[name='simpleSelection']:checked").val(),
	DisplayMode: $("input[name='displayMode']:checked").val(),
	LayerTree: $("#tree-layer").val(),
	LayerCirculation: $("#circulation-layer").val(),
	LayerSiteBuilding: $("#sitebuilding-layer").val(),
	LayerProjectBuilding: $("#projectbuilding-layer").val(),
	LayerSunShadow: $("#sunshadow-layer").val(),
	LayerRoad: $("#road-layer").val()
    }
}
