<?php

namespace App\Http\Controllers;

use App\Models\User;
use App\Models\Users\UserData;
use Carbon\Carbon;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use App\Notifications\SignupActivate;
use Laravolt\Avatar\Avatar;
use Illuminate\Support\Facades\Storage;

class AuthController extends Controller
{
    /**
     * Login user and create token
     *
     * @param Request $request
     * @return JsonResponse
     */

    public function login(Request $request)
    {
        $request->validate([
            'email' => 'required|string|email',
            'password' => 'required|string'
        ]);
        $credentials = request(['email', 'password']);
        if (!Auth::attempt($credentials))
            return response()->json([
                'error' => 'Email or Password is wrong'
            ], 401);
        $user = $request->user();
        if ($user->locked) {
            return response()->json([
                'error' => 'User has been locked'
            ], 401);
        }
        $tokenResult = $user->createToken('Personal Access Token');
        $token = $tokenResult->token;
        $token->save();
        return response()->json([
            'id' => $user->id,
            'access_token' => $tokenResult->accessToken,
            'role' => $user->role,
            'email' => $user->email,
            'token_type' => 'Bearer',
            'about' => $user->about,
            'avatar' => $user->avatar,
            'temp_password' => $user->temp_password,
            'name' => $user->name,
            'lastname' => $user->lastname,
            'created_at' => $user->created_at,
            'updated_at' => $user->updated_at,
            'expires_at' => Carbon::parse(
                $tokenResult->token->expires_at
            )->toDateTimeString()
        ]);

    }

    /**
     * Logout user (Revoke the token)
     *
     * @param Request $request
     * @return JsonResponse
     */
    public function logout(Request $request)
    {
        $request->user()->token()->revoke();
        return response()->json([
            'message' => 'Successfully logged out'
        ]);
    }

    /**
     * Signs up the new user
     *
     * @param Request $request
     * @return JsonResponse
     */

    public function signup(Request $request)
    {
        $request->validate([
            'name' => 'required|string',
            'lastname' => 'required|string',
            'email' => 'required|string|email'
        ]);
        try {
            $password = substr(str_shuffle(str_repeat($x = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', ceil(12 / strlen($x)))), 1, 12);
            $user = new User([
                'email' => $request->email,
                'password' => bcrypt($password),
                'temp_password' => $password,
                'role' => $request->role,
                'name' => $request->name,
                'lastname' => $request->lastname,
            ]);
            $user->save();
        } catch (\Exception $error){
            return response()->json([
                'message' => 'User email ' . $request->email . ' already taken!'
            ], 422);
        }
        $user->notify(new SignupActivate($user));
        $avatar = (new Avatar)->create($user->name . " " . $user->lastname)->getImageObject()->encode('png');
        Storage::put('public/avatars/'.$user->id.'/avatar.png', (string) $avatar);
        return response()->json([
            'message' => 'User ' . $request->email . ' Created!',
        ], 201);
    }

    public function reset(Request $request){
        $password = substr(str_shuffle(str_repeat($x = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', ceil(12 / strlen($x)))), 1, 12);
        $user = User::find($request->id);
        $user->update([
            'temp_password' => $password,
            'password' => bcrypt($password)]);
        $user->notify(new SignupActivate($user));
        return response()->json([
            'message' => 'Temp password created and sent',
        ], 201);
    }

    public function lock(Request $request){
        $user = User::find($request->id);
        if ($user->locked) {
            $user->update([
                'locked' => 0]);
            return response()->json([
                'message' => 'The account is now unlocked',
            ], 201);
        } else{
            $user->update([
                'locked' => 1]);
            return response()->json([
                'message' => 'The account is now Locked',
            ], 201);
        }
    }
}
