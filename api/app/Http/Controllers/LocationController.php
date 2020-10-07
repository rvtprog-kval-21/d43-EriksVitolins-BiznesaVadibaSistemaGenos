<?php

namespace App\Http\Controllers;

use App\Http\Requests\LocationNew;
use App\Models\location\Location;
use Illuminate\Http\Request;

class LocationController extends Controller
{
    public function new(LocationNew $request){

        $location = new Location([
            'Name' => $request->name,
        ]);
        $location->save();
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }

    public function index(){
        return response()->json([
            'data' => Location::all(),
        ], 201);
    }

    public function update(LocationNew $request){
        $location = Location::find($request->id);
        $location->Name = $request->name;
        $location->save();
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }

    public function delete(Request $request){
        Location::destroy($request->id);
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }
}
