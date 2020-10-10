<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Storage;

class UsersController extends Controller
{
    public function index(Request $request)
    {
        $data = $request->sort_by;
        $data = explode(' ',$data);
        $users = DB::table('users')->select('id','avatar','email','role','created_at')->orderBy($data[0],$data[1])->paginate(10);
        return response()->json([
            'data' => $users,
            'test' => $data
        ], 201);
    }

    public function search(Request $request)
    {
        $data = $request->search;
        $users = DB::table('users')->where('email', 'like', "%$data%")
            ->get();

        if (count($users) > 0){
            return response()->json([
                'data' => ['data' => $users]
            ], 201);
        }
        else{
            return response()->json([
                'errors' => "No such user found!"
            ], 201);
        }
    }

    public function user($id = null) {
        $user = User::find($id);
        return response()->json([
            'data' => $user,
        ], 201);
    }
}
