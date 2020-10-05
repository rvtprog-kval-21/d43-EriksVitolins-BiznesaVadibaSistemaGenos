<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class CheckIfAdmin
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle(Request $request, Closure $next)
    {
        $userRole = Auth::user()->role;
        if ($userRole == 'admin') {
            return $next($request);
        } else {
            return response()->json([
                'error' => 'user is not authorized to make such changes',
            ], 401);
        }
    }
}
