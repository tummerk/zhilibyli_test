<?php

namespace App\Http\Controllers;

use App\Contracts\Services\PostServiceInterface;
use App\Http\Requests\StorePostRequest;
use App\Http\Requests\UpdatePostRequest;
use App\Http\Resources\PostResource;
use App\Services\PostService;
use Illuminate\Http\JsonResponse;

class PostController extends Controller
{
    public function __construct(private PostServiceInterface $postService){}

    /**
     * Display a listing of the resource.
     */
    public function index(): JsonResponse
    {
        return $this->handleServiceCall(
            fn() => PostResource::collection($this->postService->getAllPosts()),
            'index');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StorePostRequest $request): JsonResponse
    {
        return $this->handleServiceCall(
            fn() => new PostResource($this->postService->createPost($request->validated())),
            'store',
            'Post created successfully',
            201,

        );
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id): JsonResponse
    {
        return $this->handleServiceCall(
            fn() => new PostResource($this->postService->getPostById($id)),
            'show',
            'Post show successfully',
            200
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdatePostRequest $request, string $id): JsonResponse
    {
        return $this->handleServiceCall(
            fn() => new PostResource($this->postService->updatePost($id, $request->validated())),
            'update',
            'Post updated successfully',
            200,
            );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id): JsonResponse
    {
        return $this->handleServiceCall(
            function () use ($id) {
                $this->postService->deletePost($id);
                return null;
            },
            'delete',
            'Post deleted successfully'
        );
    }

    public function count(): JsonResponse
    {
        return $this->handleServiceCall(
            fn() => ['total' => $this->postService->getPostsCount()],
            'count',
            'posts counting successfully',
            200
        );
    }
}
