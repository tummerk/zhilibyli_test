<?php

namespace App\Contracts\Services;

use App\Models\Post;
use Illuminate\Pagination\LengthAwarePaginator;

interface PostServiceInterface
{
    public function getAllPosts(): LengthAwarePaginator;

    public function getPostById(string $id): Post;

    public function createPost(array $data): Post;

    public function updatePost(string $id, array $data): Post;

    public function deletePost(string $id): void;

    public function getPostsCount(): int;

}
