<?php

use App\Http\Controllers\LocationController;
use App\Http\Controllers\UsersController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::group([
    'prefix' => 'auth'
], function () {
    Route::post('login', 'App\Http\Controllers\AuthController@login');
    Route::group([
        'middleware' => 'auth:api'
    ], function() {
        Route::get('logout', 'App\Http\Controllers\AuthController@logout');
       // Route::get('user', 'AuthController@user');
    });
});

Route::middleware(['auth:api','isAdmin'])->group(function (){
    Route::post('signup', 'App\Http\Controllers\AuthController@signup');
    Route::post('users/search', 'App\Http\Controllers\UsersController@search');
    Route::post('users', 'App\Http\Controllers\UsersController@index');
    Route::post('auth/user/passreset','App\Http\Controllers\AuthController@reset');
    Route::post('auth/user/lock','App\Http\Controllers\AuthController@lock');
    Route::post('locationAdd',[LocationController::class,'new']);
    Route::post('locations/edit',[LocationController::class,'update']);
    Route::post('locations/delete',[LocationController::class,'delete']);
});

Route::middleware(['auth:api'])->group(function (){
    Route::get('user/{id}','App\Http\Controllers\UsersController@user');
    Route::get('locations',[LocationController::class,'index']);
});
