<?php

namespace App\Http\Controllers;

use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\Log;
use Throwable;

abstract class Controller
{
    protected function success($data, $message = null, $code = 200) // для единообразия ответов
    {
        return response()->json([
            'success' => true,
            'data' => $data,
            'message' => $message
        ], $code);
    }

    protected function error($message, $code = 404)
    {
        return response()->json([
            'success' => false,
            'data' => null,
            'message' => $message
        ], $code);
    }

    protected function handleServiceCall(
        callable $serviceCall,
        string $errorContext = '',
        ?string $successMessage = null,
        int $successCode = 200
    ): JsonResponse {
        try {
            $result = $serviceCall();
            return $this->success($result, $successMessage, $successCode);
        } catch (ModelNotFoundException $e) {
            return $this->error('Resource not found.', 404);
        } catch (Throwable $e) {
            Log::error("Failed to {$errorContext}", ['exception' => $e]);
            return $this->error("Failed to {$errorContext} post", 500);
        }
    }

}
