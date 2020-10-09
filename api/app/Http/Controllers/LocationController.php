<?php

namespace App\Http\Controllers;

use App\Http\Requests\location\LocationNew;
use App\Http\Requests\location\LocationUser;
use App\Models\Location;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use function GuzzleHttp\Promise\all;

class LocationController extends Controller
{
    public function new(LocationNew $request)
    {

        $location = new Location([
            'Name' => $request->name,
        ]);
        $location->save();
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }

    public function index()
    {
        $locations = Location::all();
        foreach( $locations as $iter){
            $iter->count = DB::table('location_user')->where('location_id', $iter->id)->count();
        }
        return response()->json([
            'data' => $locations
        ], 201);
    }

    public function update(LocationNew $request)
    {
        $location = Location::find($request->id);
        $location->Name = $request->name;
        $location->save();
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }

    public function delete(Request $request)
    {
        Location::destroy($request->id);
        return response()->json([
            'message' => 'Location Saved',
        ], 201);
    }

    public function join(LocationUser $locationUser)
    {
        $location = Location::find($locationUser->location);
        $location->users()->attach(User::find($locationUser->user));
        return response()->json([
            'data' => $location,
        ], 201);
    }

    public function users($id = null)
    {
        $location = Location::find($id);
        return response()->json([
            'data' => $location->users,
        ], 201);
    }

    public function leave(LocationUser $locationUser) {
        DB::table('location_user')->where('user_id', $locationUser->user)->delete();
        return response()->json([
            'message' => 'User left location',
        ], 201);
    }
}
