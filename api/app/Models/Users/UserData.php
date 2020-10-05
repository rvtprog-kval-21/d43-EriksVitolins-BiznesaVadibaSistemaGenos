<?php

namespace App\Models\Users;

use App\Models\User;
use Illuminate\Database\Eloquent\Model;

class UserData extends Model
{
    protected $fillable = [
        'user_id','name', 'lastname', 'title', 'about'
    ];

    public function user() {
        return $this->belongsTo(User::class);
    }
}
