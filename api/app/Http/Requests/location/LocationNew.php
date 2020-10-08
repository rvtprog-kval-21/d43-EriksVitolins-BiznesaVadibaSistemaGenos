<?php

namespace App\Http\Requests\location;

use Illuminate\Foundation\Http\FormRequest;

class LocationNew extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'name' => 'required|unique:locations|min:2'
        ];
    }

    public function messages()
    {
        return [
            'name.required' => 'A name is required',
            'name.unique' => 'The name isn\'t unique',
            'name.min' => 'The name need to be atleast one character'
        ];
    }
}
